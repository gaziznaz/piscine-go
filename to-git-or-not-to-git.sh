#! /bin/bash
id=$(curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"nazerkegazizova\"}}){id}}"}')
echo $id | jq .[].user[0].id
