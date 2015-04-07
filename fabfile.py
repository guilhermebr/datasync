from fabric.api import env, execute, get, put, task, run, local, roles


env.roledefs = {
    'local': ['127.0.0.1'],
}

@task
def run_cassandra():
    cassandra_image = "spotify/cassandra"
    cassandra_params = "cassandra -p 7199:7199 -p 9160:9160 "
    docker_run("-dti --name " + cassandra_params + cassandra_image, True)


# Docker Commands
def docker(cmd, is_local=False):
    if is_local:
        return local("docker %s" % cmd)
    return run("docker %s" % cmd)


def docker_start(container, local=None):
    return docker("start %s" % container, local)


def docker_stop(container, local=None):
    return docker("kill %s" % container, local)


def docker_clean(container, local=None):
    docker_stop(container)
    return docker("rm %s" % container, local)


def docker_restart(container, local=None):
    docker_stop(container, local)
    return docker_start(container, local)


def docker_pull(image, local=None):
    return docker("pull %s" % image, local)


def docker_run(command, local=None):
    return docker("run %s" % command, local)


def docker_build(container, name, local=None):
    return docker("build -t %s %s " % (name, container), local)
