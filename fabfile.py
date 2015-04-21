from fabric.api import env, execute, get, put, task, run, local, roles
import platform

env.roledefs = {
    'local': ['127.0.0.1'],
}

@task
def run_cassandra():
    cassandra_image = "spotify/cassandra"
    cassandra_params = "--name cassandra -p 7199:7199 -p 9160:9160 -p 9042:9042"

    try:
        docker_restart('cassandra', True)
    except:
        docker_run("-dti " + cassandra_params + cassandra_image, True)

@task
def run_elastic():
    es_image = "elasticsearch"
    es_params = "--name elasticsearch -p 9200:9200 -p 9300:9300"

    try:
        docker_restart('elasticsearch', True)
    except:
        docker_run("-dti " + es_params + " " + es_image, True)

@task
def init_cassandra():
    with docker_exec('cassandra', True):
        run('cqlsh')
        local("create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };")
        local('create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));')
        local('create index on example.tweet(timeline);')
        local('exit')

#
# def prepare_docker():
#     if platform.system() == "Darwin"
#

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
    docker_start(container, local)


def docker_pull(image, local=None):
    return docker("pull %s" % image, local)


def docker_run(command, local=None):
    return docker("run %s" % command, local)


def docker_build(container, name, local=None):
    return docker("build -t %s %s " % (name, container), local)


def docker_exec(container, local=None):
    docker('exec -ti %s bash' % container, local)
