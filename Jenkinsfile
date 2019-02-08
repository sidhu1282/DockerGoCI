//Creating Jenkins Pipeline
pipeline {
    agent any
    //Initializing Source paths
environment {
projectName = 'DockerGoCI'
fullVersion = "${BUILD_NUMBER}"
dockerImage = "${projectName}_${BUILD_NUMBER}.tar"
artifactoryUrl = "https://artifactoryurl/artifactory"
artifactoryTarget = 'AABFH-GENERIC-SNAPSHOT/docker/DockerGoCI'
ARTIFACTORY = credentials('ARTIFACTORY-ENCRYPTED-PWD')

}
    //Using maven tool to build code
    tools {
        maven 'Maven 3.3.9'

    }
    // Maven build stages
    stages {
        stage ('Initialize') {
            steps {
                sh '''
                    echo "PATH = ${PATH}"
                    echo "M2_HOME = ${M2_HOME}"
                '''
            }
        }
        //Git Checkout step  
stage('Checkout') {
    steps {
deleteDir()
git url: 'https://github.com/sidhu1282/DockerGoCI', credentialsId: 'GIT-TOKEN', branch: 'master'
}
    // Trigger Unit TestCases
    stage("test-go") {
      agent {
        label "go"
      }
      steps {
        sh "go get -d -v -t && go test --cover -v ./... --run UnitTest && go build -v -o go-demo"
      }
    }
    stage("test-docker") {
      steps {
        sh "docker container run -v ${workspace}:/usr/src/myapp -w /usr/src/myapp golang:1.9 bash -c \"go get -d -v -t && go test --cover -v ./... --run UnitTest && go build -v -o go-demo\""
      }
    }
    stage("test-dc") {
      steps {
        sh "docker-compose run --rm unit"
      }
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
//build docker image and set latest commit number in lastcommitsha ,  set application version parameter from jenkins followed by buildnumber
        stage('Build docker image') {
     steps {
sh "rm -f $projectName*.tar"
         sh "docker build -f DockerGoCI/Docker . -t $projectName:$fullVersion"
         sh "docker save $projectName:$fullVersion > $dockerImage"
         sh "ls -lh $dockerImage"
     }
 }
            //push the image into JFrog Artifactory
    stage('Artifactory upload'){
steps {
                    sh "jfrog rt u --user $ARTIFACTORY_USR --password $ARTIFACTORY_PSW --url $artifactoryUrl $dockerImage $artifactoryTarget/$projectName/"
            }
        }
      }
    }
	//Deploy the Docker image locally and publish on 8081 host port.
       stage('Deploy'){
         echo 'ssh to web server and tell it to pull new image'
	  sh 'docker rm -f $(docker ps -aq)'
	       sh 'docker run --name=nodejs -d -p 8081:8080 sidhu1282/DockerGoCI:${version}_${BUILD_NUMBER}'
       }
    }	
    catch (err) {
	//Set the build result to failure is some exception occur.
        currentBuild.result = "FAILURE"
        throw err
    }
}
