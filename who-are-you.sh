echo "You just woke up in a dark alley... You can not remember who you are... The only though that comes to your mind is a tag that says:"
curl  https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq ' . [] | select(.id==70)|.name'


