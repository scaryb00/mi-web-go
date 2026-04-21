pipeline {
    agent any
    environment {
        // Asegúrate de que este nombre sea el que usas en el deployment.yaml
        IMAGE_NAME = 'mi-web-uoc' 
        IMAGE_TAG = 'latest'
    }
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/scaryb00/mi-web-go.git'
            }
        }
        stage('Build Docker Image') {
            steps {
                // CORRECCIÓN: Comillas dobles para que las variables funcionen
                sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG} ."
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s/deployment.yaml'
                sh 'kubectl apply -f k8s/service.yaml'
                sh 'kubectl rollout restart deployment/mi-web-deployment'
            }
        }
    }
    post {
        success { echo '¡Despliegue completado con éxito!' }
        failure { echo 'El pipeline ha fallado. Revisa los logs.' }
    }
}
