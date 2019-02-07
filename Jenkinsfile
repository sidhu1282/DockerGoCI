pipeline {
    agent any
environment {
projectName = 'DockerGoCI'
fullVersion = "${BUILD_NUMBER}"
dockerImage = "${projectName}_${BUILD_NUMBER}.tar"
artifactoryUrl = "https://artifactoryurl/artifactory"
artifactoryTarget = 'AABFH-GENERIC-SNAPSHOT/docker/DockerGoCI'
ARTIFACTORY = credentials('ARTIFACTORY-ENCRYPTED-PWD')

}
    tools {
        maven 'Maven 3.3.9'

    }
    stages {
        stage ('Initialize') {
            steps {
                sh '''
                    echo "PATH = ${PATH}"
                    echo "M2_HOME = ${M2_HOME}"
                '''
            }
        }
stage('Checkout') {
    steps {
deleteDir()
git url: 'https://github.com/sidhu1282/DockerGoCI', credentialsId: 'GIT-TOKEN', branch: 'master'
}
post {
                success {
                    junit 'target/surefire-reports/**/*.xml'
                }
            }
 }
        stage ('Build') {
            steps {
                sh 'mvn -Dmaven.test.failure.ignore=true install'
            }

        stage('Build docker image') {
     steps {
sh "rm -f $projectName*.tar"
         sh "docker build -f DockerGoCI/Docker . -t $projectName:$fullVersion"
         sh "docker save $projectName:$fullVersion > $dockerImage"
         sh "ls -lh $dockerImage"
     }
 }
    stage('Artifactory upload'){
steps {
                    sh "jfrog rt u --user $ARTIFACTORY_USR --password $ARTIFACTORY_PSW --url $artifactoryUrl $dockerImage $artifactoryTarget/$projectName/"
            }
        }
      }
    }
}
