curl https://01.gritlab.ax/assets/superhero/all.json | jq ' .[] | select( .id == '$HERO_ID' ) | .connections | .relatives' | tr -d "\'" | tr -d "\""
