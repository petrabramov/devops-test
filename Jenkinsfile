#!/usr/bin/groovy

node {
    def app
    def gitBranch = 'master'
    def gitUrl = 'https://github.com/petrabramov/devops-test'
    def dockerImage = 'petrabramov/devops-test'
    def hubCredentialID = 'docker-hub-credentials'
    def containerName = 'devops-test'

    stage('Checkout SCM') {
        try {
            checkout([
                $class: 'GitSCM',
                branches: [[name: gitBranch]],
                userRemoteConfigs: [[url: gitUrl]]
            ])
        
            slackSend(message: "Checkout successful")
        }
        catch (Exception e) {
            slackSend(color: "danger", message: "Checkout failed \n $e")
            error "Checkout failed"
        }
    }

    stage('Build app') {
        try {
            app = docker.build(dockerImage)

            slackSend(message: "Build successful")
        }
        catch (Exception e) {
            slackSend(color: "danger", message: "Build failed \n $e")
            error "Build failed"
        }
    }

    stage('Test app') {
        try {
            result = sh(
                script: "docker run --rm -v /var/lib/jenkins/workspace/build/src:/app -w /app golang:latest go test -v",
                returnStdout: true
            ).trim()

            if (result.contains("FAIL")) {
                throw new Exception(result)
            }

            slackSend(message: "Test successful \n $result")
        }
        catch (Exception e) {
            slackSend(color: "danger", message: "Test failed \n $e")
            error "Test failed"
        }
    }

    stage('Code analyze') {
        try {
            result = sh(
                script: "docker run --rm -v /var/lib/jenkins/workspace/build/src:/app -w /app golang:latest /bin/bash -c 'go get honnef.co/go/tools/cmd/staticcheck; go install honnef.co/go/tools/cmd/staticcheck; staticcheck .'",
                returnStdout: true
            )

            if (result) {
                throw new Exception(result)
            }

            slackSend(message: "Code analyze successful")
        }
        catch (Exception e) {
            slackSend(color: "danger", message: "Code analyze failed \n $e")
            error "Code analyze failed"
        }
    }

    stage('Push image') {
        try {
            docker.withRegistry('https://registry.hub.docker.com', hubCredentialID) {
                app.push("${env.BUILD_NUMBER}")
                app.push("latest")
            }

            slackSend(message: "Publish successful")
        }
        catch (Exception e) {
            slackSend(color: "danger", message: "Push failed \n $e")
            error "Push failed"
        }
    }

    stage('Deployment') {
        def message

        try {
            sh "docker pull $dockerImage"
            sh "docker stop $containerName && docker rm $containerName"
            sh "docker run -d -p 8000:80 --name=$containerName $dockerImage"

            message = sh(
                script: "docker inspect --format '{{ (.State.Status) }}' $containerName",
                returnStdout: true
            ).trim()

            if (message == "running") {
                slackSend(color: "good", message: "Deploy successful \n $message")
            } else {
                throw new Exception("container not running")
            }
        }
        catch (Exception e) {
            slackSend(color: "danger", message: "Deployment failed \n $e")
            error "Deployment failed" 
        }
    }
}