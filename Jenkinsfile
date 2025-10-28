pipeline{

    agent any

    environment{
        IMAGE_NAME="linebot"
    }

    stages{
        stage("Build and Push"){
            steps{
                echo 'building....'
                script{
                def version = "v1.${env.BUILD_NUMBER}"
                def registry = sh(script: 'getent hosts host.docker.internal | awk \'{print $1}\' || true', returnStdout: true).trim()
                sh """
                    docker buildx build --platform linux/arm64 -t ${registry}:5000/${IMAGE_NAME}:${version} --load .
                    docker push ${registry}:5000/${IMAGE_NAME}:${version}
                """
            }
            }
            
            
        }
    }
}