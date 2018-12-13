rm -f metrics.json
rm -f plot.html
rm -f results.bin
echo "GET http://localhost:6767/accounts/10001" | vegeta attack -duration=5s -rate 1000 | tee results.bin | vegeta report
vegeta report -type=json results.bin > metrics.json
cat results.bin | vegeta plot > plot.html
open -a 'Google Chrome' plot.html