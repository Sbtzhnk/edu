#!/bin/bash
word=""
hint=""
wordMask=""
isWordCompleted=0
playerName=""
tryCount=0
#check commit
function ToLower {
    echo "${1,,}"
}

function CheckAnswer() {
    local answer="$1"
    local isSymbolRepeated=0
    local isSymbolGuessed=0
    local hiddenSymbolsCount=0
    local answerLower=$(ToLower  "$answer")
    local wordLower=$(ToLower "$word")
    
    if [ ${#answer} -eq ${#word} ]; then
        if [ "$answerLower" = "$wordLower" ]; then
            echo "Слово успешно отгадано!"
            isWordCompleted=1
            return 0
        else
            echo "Неверное слово! Попробуйте еще раз."
            return 1
        fi
    else
        if [ ${#answer} -eq 1 ]; then
            local charLower=$(ToLower "$answer")
            
            for (( wordMaskIdx=0; wordMaskIdx<${#wordMask}; wordMaskIdx++ )); do
                local maskChar="${wordMask:$wordMaskIdx:1}"
                local maskCharLower=$(ToLower "$maskChar")
                if [ "$charLower" = "$maskCharLower" ]; then
                    isSymbolRepeated=1
                    break
                fi
            done
            
            if [ $isSymbolRepeated -eq 1 ]; then
                echo "Эта буква '$answer' уже была угадана ранее!"
                return 1
            fi
            
            local newWordMask=""
            isSymbolGuessed=0
            
            for (( wordIdx=0; wordIdx<${#word}; wordIdx++ )); do
                local wordChar="${word:$wordIdx:1}"
                local wordCharLower=$(ToLower "$wordChar")
                
                if [ "$charLower" = "$wordCharLower" ]; then
                    newWordMask+="$wordChar"
                    isSymbolGuessed=1
                else
                    newWordMask+="${wordMask:$wordIdx:1}"
                fi
            done
            
            wordMask="$newWordMask"
            
            hiddenSymbolsCount=0
            for (( i=0; i<${#wordMask}; i++ )); do
                if [ "${wordMask:$i:1}" = "#" ]; then
                    hiddenSymbolsCount=$((hiddenSymbolsCount + 1))
                fi
            done
            
            if [ $isSymbolGuessed -eq 1 ]; then
                echo "Буква '$answer' есть в слове!"
                echo "Текущее состояние: $wordMask"
            else
                echo "Буквы '$answer' нет в слове! Попробуйте снова."
                return 1
            fi
            
            if [ $hiddenSymbolsCount -eq 0 ]; then
                echo "Поздравляем! Слово '$word' полностью отгадано!"
                isWordCompleted=1
                return 0
            fi
            return 0
        else
            echo "Введите одну букву или слово целиком (длина слова: ${#word})"
            return 0
        fi
    fi
}

function GetWordAndHints() {
    hintsArray=(
        "Это сладкий красный фрукт"
        "Это домашнее животное, которое мяукает"
        "Этот транспорт ездит по рельсам"
        "Этот предмет используется для письма"
        "Это небесное тело, светящее ночью"
    )

    wordsArray=(
        "яблоко"
        "кошка"
        "поезд"
        "ручка"
        "луна"
    )

    randId=$((RANDOM % ${#hintsArray[@]}))

    hint="${hintsArray[$randId]}"
    word="${wordsArray[$randId]}"
    
    wordMask=""
    for (( maskIdx=0; maskIdx<${#word}; maskIdx++ )); do
        wordMask+="#" 
    done
}

echo "Добро пожаловать в игру \"Поле чудес\"."
sleep 0.5
echo "Представься, игрок!"
sleep 0.5
read -p "Введите имя: " playerName
sleep 0.5
echo "Привет, $playerName!"
sleep 0.5

GetWordAndHints

echo "Тебе нужно отгадать слово."
sleep 0.5
echo "Подсказка: $hint"
sleep 0.5
echo "Слово: $wordMask"
sleep 0.5 
echo "Введи букву или отгадай сразу слово целиком, $playerName."

while [ $isWordCompleted -eq 0 ]; do
    if [ $tryCount -gt 9 ]; then
        break
    fi
    
    sleep 0.5
    read -p "Введите ответ: " answer
    sleep 0.5

    if [ -z "$answer" ]; then
        echo "Вы ничего не ввели. Попробуйте снова."
        tryCount=$((tryCount+1))
        continue
    fi
    
    if [ ${#answer} -ne ${#word} ] && [ ${#answer} -ne 1 ]; then
        echo "Некорректный ответ! Введите одну букву или слово целиком (длина слова: ${#word})."
        tryCount=$((tryCount+1))
        continue
    fi
    
    CheckAnswer "$answer" 
    if [ $? -eq 1 ]; then
        tryCount=$((tryCount+1))
    fi
    
    if [ $isWordCompleted -eq 1 ]; then
        break
    fi
done

sleep 0.5
if [ $tryCount -gt 9 ]; then
    echo "$playerName" проиграл. Начни игру заново! 
else 
    echo "Игра завершена! Спасибо за игру, $playerName!"
fi
