range=$1
case=$2
declare -i counter=$range
numbers=""
numbersArr=()

for ((i = 1; i <= $range; i++)); do
  declare -i numberToAppend="$((RANDOM % 1000))"
  repeats=no
  while [[ "$repeats" = "no" ]]; do
    repeats=yes
    for y in "${numbersArr[@]}"; do
      if [ "$y" = "$numberToAppend" ]; then
        repeats=no
        numberToAppend=$(($y + 1))
      fi
    done
  done
  numbers+="$numberToAppend"
  numbersArr+=("$numberToAppend")
  if [ "$i" != "$counter" ]; then
    numbers+=" "
  fi
done

if [ "$case" = 0 ]; then
  ./push-swap "$numbers"
  echo -n "number of instructions: "
  ./push-swap "$numbers" | wc -l
else
  ./push-swap "$numbers" | ./checker "$numbers"
fi
