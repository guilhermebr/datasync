from fabric.api import env, execute, get, put, task, local


def run_cassandra():
    cassandra_image = "spotify/cassandra"
    cassandra_params = "cassandra -p 7199:7199 -p 9160:9160 "
    docker_run("-dti --name " + cassandra_params + cassandra_image)


# Docker Commands
def docker(cmd):
    return local("docker %s" % cmd)


def docker_start(container):
    return docker("start %s" % container)


def docker_stop(container):
    return docker("kill %s" % container)


def docker_clean(container):
    docker_stop(container)
    return docker("rm %s" % container)


def docker_restart(container):
    docker_stop(container)
    return docker_start(container)


def docker_pull(image):
    return docker("pull %s" % image)


def docker_run(command):
    return docker("run %s" % command)


def docker_build(container, name):
    return docker("build -t %s %s " % (name, container))
