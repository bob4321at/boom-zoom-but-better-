var getUsername = document.getElementById("username");
var getNewRoomName = document.getElementById("new roomname");
var getExistingRoomName = document.getElementById("existing roomname");

function makeRoom() {
  let name = getNewRoomName.value;
  sessionStorage.setItem("lobby", name)
  fetch("/makeRoom", {
    "method": "POST",
    "body": JSON.stringify({"Name": name, "Users": 1}),
  })
  window.location.href = "/game"
}

async function joinRoom() {
  let name = getExistingRoomName.value;

  let room = await fetch("/getRoom?name=" + name, {
    "method": "GET",
  })
  let json = await room.json()

  if (json.Error == "incorrect lobby name") {
    alert("invalid lobby name")
    return
  }

  if (json.Users > 1) {
    alert("lobby full")
    return
  }
  console.log(json.Users)

  sessionStorage.setItem("lobby", name);

  fetch("/addUserToRoom?name="+name, {
    "method": "GET"
  })

  window.location.href = "/game"
}
