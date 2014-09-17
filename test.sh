curl -X GET http://localhost:3000/game 
curl -X POST http://localhost:3000/game/koan --data "{\"koan\":\"1^SG\"}" 
curl -X POST http://localhost:3000/game/guess --data "{\"rule\":\"1^\"}" 
curl -X POST http://localhost:3000/game --data "{\"rule\": \"2^\", \"true\":\"3^G\", \"false\":\"1^G\"}" 
curl -X POST http://localhost:3000/game --data "{\"rule\": \"2^\", \"true\":\"3^G\", \"false\":\"1^G\"}" 
curl -X POST http://localhost:3000/game/koan --data "{\"koan\":\"1^SG\"}" 
curl -X POST http://localhost:3000/game/guess --data "{\"rule\":\"2^\"}" 
curl -X POST http://localhost:3000/game/koan --data "{\"koan\":\"1^SG\"}" 
curl -X POST http://localhost:3000/game/generate 
curl -X POST http://localhost:3000/game/koan --data "{\"koan\":\"1^SG\"}" 
curl -X POST http://localhost:3000/game/koan --data "{\"koan\":\"3^SG\"}" 
curl -X POST http://localhost:3000/game/generate 
curl -X POST http://localhost:3000/game/koan --data "{\"koan\":\"1^SG\"}" 