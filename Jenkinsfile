pipeline{

    agent any

    environment{
        IMAGE_NAME="linebot"
    }

    stages{
        stage('Set Env'){
            steps{
                script{
                    env.VERSION = "v1.${env.BUILD_NUMBER}"
                    env.REGISTRY = sh(script: 'getent hosts host.docker.internal | awk \'{print $1}\' || true', returnStdout: true).trim()
                }
            }
        }
        stage("Build Image"){
            steps{
                echo 'building....'
                script{
                sh """
                    docker buildx build --platform linux/arm64 -t ${env.REGISTRY}:5000/${IMAGE_NAME}:${env.VERSION} --load .
                """
                }
            }
        }
        stage('Push Image'){
            steps{
                echo 'pushing'
                script{
                    sh"""
                    docker push ${env.REGISTRY}:5000/${IMAGE_NAME}:${env.VERSION}
                    """
                }
            }
        }
    }
}