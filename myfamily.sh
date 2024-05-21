#! /bin/bash
curl -s 'https://01.gritlab.ax/assets/superhero/all.json' | jq -c --arg hero_id "$HERO_ID" 'map(select(.id == ($hero_id|tonumber)))' | jq -c '.[] | .connections.relatives' | tr -d '"'