#!/bin/bash
echo "Enter your age:";
read age

if [ $age -le 1 ]
then 
    echo "Вы ещё не родились"
elif [ $age -ge 1 ] && [ $age -le 7 ]; 
then 
    echo "Вы ребёнок"
elif [ $age -ge 7 ]  && [ $age -le 18 ];
then
    echo "Вы школьник"
elif [ $age -ge 18 ] && [ $age -le 25 ];
then
    echo "Вы студент"
elif [ $age -ge 26] && [ $age -le 60 ];
then
    echo "Вы инжинер"
elif [ $age -ge 60 ] && [ $age -le 100 ];
then
    echo "Вы пенсионер"
elif [ $age -ge 100 ];
then
    echo "Вы долгожитель"
fi
