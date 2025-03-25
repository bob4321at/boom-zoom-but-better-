let lobby = sessionStorage.getItem("lobby")

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
