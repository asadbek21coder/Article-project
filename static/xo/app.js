var scoreX = 0
var score0 = 0
var count = 0
isEnable = [true, true, true, true, true, true, true, true, true]
simbols = ['', '', '', '', '', '', '', '', '']
document.getElementById("scoreX").innerText="Score X : "+scoreX;
document.getElementById("score0").innerText="Score O : "+score0;

function reload() {
    count = 0
        for (let i = 1; i < 10; i++) {
        document.getElementById("button" + i).innerText = "";
        simbols[i - 1] = '';
        isEnable[i-1]=true;
        document.getElementById("result").innerText= ""//"Game is continuing"

    }
document.getElementById("scoreX").innerText="Score X : "+scoreX;
document.getElementById("score0").innerText="Score O : "+score0;
}
function stop() {
    for (let i = 0;i<9;i++){
        isEnable[i]=false;
    }
}
 function addScore(simbol){
     if (simbol==="O")
     {score0++}
     else {scoreX++}
 }

function check() {
    var  text1='';
    // check horizontally

    for (let i = 0; i < 7; i = i + 3) {
         text1 = '';
        for (let j = i; j < i + 3; j++) {
            text1 += simbols[j]
        }
        if (text1 == "OOO" || text1 == "XXX") {
            document.getElementById("result").innerText = text1[1] + " wins horizontally"
            stop();
            addScore(text1[1]);
            break;
        }
    }
    // check vertically
    for (let i = 0; i < 3; i++) {
         text1 = '';
        for (let j = i; j < i + 7; j=j+3) {
            text1 += simbols[j];
        }
        if (text1 == "OOO" || text1 == "XXX") {
            document.getElementById("result").innerText = text1[1] + " wins vertically"
            addScore(text1[1])
            stop();
            break;
        }
    }

    // check diagonally
    text1 = simbols[0] + simbols[4] + simbols[8];
    if (text1 == "OOO" || text1 == "XXX") {
        document.getElementById("result").innerText = text1[1] + " wins diagonally"
        addScore(text1[1])
        stop();
    }
    text1 = simbols[2] + simbols[4] + simbols[6];
    if (text1 == "OOO" || text1 == "XXX") {
        document.getElementById("result").innerText = text1[1] + " wins diagonally"
        addScore(text1[1])
        stop();
    }
}

function addAction(act) {
    if (isEnable[act - 1]) {
        if (count % 2) {
            document.getElementById("button" + act).innerText = "X"
            simbols[act - 1] = "X"
        } else {
            document.getElementById("button" + act).innerText = "O"
            simbols[act - 1] = "O"
        }
        count++
        isEnable[act - 1] = false
        if (count > 4) {
            check();
        }
    }
}
