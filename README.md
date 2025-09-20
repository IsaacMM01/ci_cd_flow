# 🚀 Proyecto DevOps: Terraform + GitHub Actions + Docker + EC2

Este proyecto muestra un flujo de Infraestructura como Código (IaC) y CI/CD usando:

- Terraform → Para aprovisionar una instancia EC2 a partir de una AMI personalizada con Docker ya instalado.
- Docker → Para construir y empaquetar una aplicación en Go.
- Docker Hub → Como registro de imágenes.
- GitHub Actions → Para automatizar el pipeline de build, push y despliegue.
- Kubernetes (futuro) → Para escalar el despliegue en un clúster.

## 1. Infraestructura con Terraform

Terraform crea:

Una instancia EC2 en AWS usando una AMI personalizada, guardada en mi cuenta de AWS basada en ubuntu y con docker ya preinstalado.

Un Security Group con acceso por SSH (22) y HTTP (80).

### Ejemplo de variables (terraform/dev.auto.tfvars):
``` t
    aws_region = "us-east-1"
    ami_id = "ami-xxxxxxxx"
    vpc_id = "vpc-xxxxxxxx"
    subnet_id = "subnet-xxxxxxxx"
    key_pair_name = "docker-app-keys"
```
### Comandos principales:
``` bash
    cd terraform/
    terraform init
    terraform plan
    terraform apply -auto-approve
```
Es necesario que al finalizar se entregue la ip pública para que pueda ser ingresada por la persona que hizo los cambios

## 2. Aplicación en Go con Docker

La aplicación de go, consta únicamente de una ruta raíz que mostrará el mensaje de "Hola desde Go" y una ruta para saluda.
El archivo Docker file es necesario para poder contruir nuestra imágen a apartir de la aplicción de go