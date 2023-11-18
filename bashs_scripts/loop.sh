
numberOfLoops=$1

if (( numberOfLoops < 100 )); then
    for ((i=1; i <= numberOfLoops; i++))
        
        do
            printf "%s times I've printed hannessoosaar\n" "$i"
        done
else 
    for ((i=1; i <= 100; i++))
        do
            printf "%s times I've printed hannessoosaar\n" "$i"
        done
ficd