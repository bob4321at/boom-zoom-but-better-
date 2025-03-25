let lobby = sessionStorage.getItem("lobby")

let messege_element = document.getElementById("messege element")
let messege_render = document.getElementById("messege list")

const timerId = window.setInterval( updateMesseges, 13 * 100 );

function sendMessege() {
  console.log(messege_element.value)
  fetch("/sendMessege?user="+sessionStorage.getItem("username")+"&messege="+messege_element.value+"&name="+lobby)
}

async function updateMesseges() {
  messege_render.innerHTML = '';

  let room = await fetch("/getRoom?name=" + lobby, {
    "method": "GET",
  })
  let json = await room.json()

  for (let mi = 0; mi < json.Messeges.length; mi++) {
    messege = document.createElement("li")
    messege.innerHTML = json.Messeges[mi].Username + ": " + json.Messeges[mi].Content
    messege_render.append(messege)
  }
}

async function getRoom() {
  let room = await fetch("/getRoom?name=" + lobby, {
    "method": "GET",
  })
  let json = await room.json()
  console.log(json)
}

window.addEventListener('beforeunload', function() {
  fetch("/removeUserFromRoom?name="+lobby, {
    "method": "GET",
  })
});

getRoom()
