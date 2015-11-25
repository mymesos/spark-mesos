package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("sbin/start-mesos-dispatcher.sh",
		"--master", os.Getenv("SPARK_MESOS_MASTER_URL"),
		"--zk", os.Getenv("SPARK_ZK_URL"),
		"--host", os.Getenv("HOST"),
		"--port", os.Getenv("PORT0"),
		"--webui-port", os.Getenv("PORT1"),
		"--name", "spark")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	for {
		if err := exec.Command("sbin/spark-daemon.sh", "status", "org.apache.spark.deploy.mesos.MesosClusterDispatcher", "1").Run(); err != nil {
			log.Fatal(err)
		}
		<-time.Tick(10 * time.Minute)
	}
}
