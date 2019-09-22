#!/usr/bin/env groovy

pipelineJob('Build_job') {
  definition {
    cps {
      sandbox()
      script("""
        node {
          def img = "alrf/go-redis-pg-demo:latest"
          try {
            stage('Checkout') {
              deleteDir()
              git(
                url: 'https://github.com/alrf/go_redis_pg.git',
                branch: 'master'
              )
            }
            stage('Build Dev') {
              sh "docker images"
              sh "docker build --target builder -t \$img -f Dockerfile ."
            }
            stage('Build Prod') {
              sh "docker images"
              sh "docker build --target app -t \$img -f Dockerfile ."
            }            
            stage('Push') {
              withDockerRegistry([credentialsId: 'temp_token', url: '']) {
                sh "docker push \$img"
              }
            }
          }
          catch(err) {
              throw(err)
          } finally {}
        }
      """.stripIndent())
    }
  }
}
