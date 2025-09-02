package testtb

import (
	"log"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/tigerbeetle/tigerbeetle-go/pkg/types"

	sdk "github.com/tigerbeetle/tigerbeetle-go"
)

func TestContainer(t *testing.T) {
	pool, errNP := dockertest.NewPool("")
	if errNP != nil {
		log.Fatalf("Could not construct pool: %s", errNP)
	}

	if errPG := pool.Client.Ping(); errPG != nil {
		log.Fatalf("Could not connect to Docker: %s", errPG)
	}

	resourceF, errF := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "ghcr.io/tigerbeetle/tigerbeetle",
		Tag:        "0.16.51",
		Entrypoint: []string{
			"/bin/sh",
			"-c",
			"mkdir /data && ./tigerbeetle format --cluster=0 --replica=0 --replica-count=1 /data/0_0.tigerbeetle && ./tigerbeetle start --addresses=0.0.0.0:3000 /data/0_0.tigerbeetle",
		},
		ExposedPorts: []string{"3000/tcp"},
		SecurityOpt:  []string{"seccomp=unconfined"},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if errF != nil {
		log.Fatalf("Could not start tigerbeetle: %s", errF)
	}

	testPort := resourceF.GetPort("3000/tcp")

	if err := pool.Retry(func() error {
		tbc, errNC := sdk.NewClient(types.ToUint128(0), []string{testPort})
		if errNC != nil {
			return errNC
		}

		if _, errLA := tbc.LookupAccounts([]types.Uint128{types.ToUint128(0)}); errLA == nil {
			tbc.Close()
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to tigerbeetle container: %s", err)
	}

	resourceF.Expire(240)

	_, errTB := sdk.NewClient(types.ToUint128(0), []string{testPort})
	if errTB != nil {
		log.Fatalf("Could not create tigerbeetle client: %s", errTB)
	}
}
