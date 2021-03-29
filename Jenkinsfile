pipeline {
    agent any
    tools {
        go 'go 1.16.2'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build /var/jenkins_home/workspace/demo2/main.go'
            }
        }

   
		stage('Test') {
			steps {
				echo 'Running vetting'
				sh 'go vet /var/jenkins_home/workspace/demo2/main.go'
				
				echo 'Running linting'
				sh 'golint /var/jenkins_home/workspace/demo2/main.go'
			}
		}
        
    }    
}