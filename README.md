## Google Cloud Run:
    
* Como usar:
  * Acesse a url:
    * https://desafio-cloud-run-81738062130.us-central1.run.app/temperatura?cep=88750000
    * Troque o cep para o desejado.
    * O retorno será um JSON com a temperatura atual do local.
* Para execução local, defina as variaveis de ambiente
  * `export WEATHER_API_KEY="chave-weather-api"`
  * `export PORT=8080`
* Para executar via docker, basta utilizar o docker-compose:
  * `docker compose up cloud-run-app`
  * Acessar em: `http://localhost:8080/temperatura?cep=88750000`
