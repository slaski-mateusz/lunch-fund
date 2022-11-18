./teams_del.sh Testy
read

./teams_new.sh Testy
read

./members_new.sh --teamname Testy --membername "Abek Abekowicz" --email "aa@testy" --phone 123456789 --isactive 1
read
./members_new.sh --teamname Testy --membername "Bebek Bekowski" --email "bb@testy" --phone 234567890 --isactive 1
read
./members_new.sh --teamname Testy --membername "Cecek Cecyk" --email "cc@testy" --phone 345678901 --isactive 1
read
./members_new.sh --teamname Testy --membername "Dedek Dedeletek" --email "dd@testy" --phone 456789012 --isactive 1
read
./members_new.sh --teamname Testy --membername "Elek Ektotester" --email "ee@testy" --phone 567890123 --isactive 1
read
./members_new.sh --teamname Testy --membername "Faka Fukajaca" --email "ff@testy" --phone 678901234 --isactive 1
read
./members_new.sh --teamname Testy --membername "Glunacja Gluś" --email "gg@testy" --phone 789012345 --isactive 1
read
./members_new.sh --teamname Testy --membername "Hajda Huper" --email "hh@testy" --phone 890123456 --isactive 1
read

./orders_new.sh --teamname Testy --ordername "Monday Pizza" --timestamp 1667827504 --deliverycost 999 --tipcost 1000
read
./orders_details_new.sh --teamname Testy --orderid 1 --memberid 1 --isfounder 1 --amount 2590
read
./orders_details_new.sh --teamname Testy --orderid 1 --memberid 2 --isfounder 0 --amount 2590
read
./orders_details_new.sh --teamname Testy --orderid 1 --memberid 3 --isfounder 0 --amount 2255
read

read
./orders_new.sh --teamname Testy --ordername "Tuesday Chinesee bar" --timestamp 1667913904 --deliverycost 999 --tipcost 1000
read
./orders_details_new.sh --teamname Testy --orderid 2 --memberid 2 --isfounder 1 --amount 3550
read
./orders_details_new.sh --teamname Testy --orderid 2 --memberid 3 --isfounder 0 --amount 2700
read
./orders_details_new.sh --teamname Testy --orderid 2 --memberid 4 --isfounder 0 --amount 2960
read

./orders_new.sh --teamname Testy --ordername "Wenesday Indian restaurant" --timestamp 1668000904 --deliverycost 0 --tipcost 2000
read
./orders_details_new.sh --teamname Testy --orderid 3 --memberid 3 --isfounder 1 --amount 3320
read
./orders_details_new.sh --teamname Testy --orderid 3 --memberid 4 --isfounder 0 --amount 3250
read
./orders_details_new.sh --teamname Testy --orderid 3 --memberid 5 --isfounder 0 --amount 3525
read

./orders_new.sh --teamname Testy --ordername "Thursday Polish Pierogi" --timestamp 1668086104 --deliverycost 499 --tipcost 1500
read
./orders_details_new.sh --teamname Testy --orderid 4 --memberid 4 --isfounder 1 --amount 2490
read
./orders_details_new.sh --teamname Testy --orderid 4 --memberid 5 --isfounder 0 --amount 2490
read
./orders_details_new.sh --teamname Testy --orderid 4 --memberid 6 --isfounder 0 --amount 2490
read

./orders_new.sh --teamname Testy --ordername "Friday Sushi" --timestamp 1668174004 --deliverycost 0 --tipcost 1000
read
./orders_details_new.sh --teamname Testy --orderid 5 --memberid 5 --isfounder 1 --amount 3050
read
./orders_details_new.sh --teamname Testy --orderid 5 --memberid 6 --isfounder 0 --amount 3150
read
./orders_details_new.sh --teamname Testy --orderid 5 --memberid 7 --isfounder 0 --amount 2970
read
