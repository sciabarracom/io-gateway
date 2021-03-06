package main

import (
	"fmt"
)

// WhiskDeploy deploys openwhisk standalone
func WhiskDeploy() error {
	fmt.Println("Deploying Whisk...")
	whiskDockerRun()
	return nil
}

// WhiskDestroy destroys openwhisk standalone
func WhiskDestroy() error {
	fmt.Println("Destroying Whisk...")
	Sys("docker exec iogw-openwhisk stop")
	fmt.Println()
	return nil
}

// return empty string if ok, otherwise the error
func whiskDockerRun() string {
	image := WhiskImage + ":" + Version
	if err := dockerPull(image); err != nil {
		return err.Error()
	}
	redisIP := dockerIP("iogw-redis")
	if redisIP == nil {
		return "cannot locate redis"
	}
	cmd := fmt.Sprintf(`docker run -d -p 3280:3280
--rm --name iogw-openwhisk --hostname openwhisk
-e CONTAINER_EXTRA_ENV=__OW_REDIS=%s
-e CONFIG_FORCE_whisk_users_guest=%s
-v //var/run/docker.sock:/var/run/docker.sock %s`,
		*redisIP, Config.WhiskAPIKey, image)
	_, err := SysErr(cmd)
	if err != nil {
		return "cannot start server: " + err.Error()
	}
	err = Run("docker exec iogw-openwhisk waitready")
	if err != nil {
		return "server readyness error: " + err.Error()
	}
	return ""
}
