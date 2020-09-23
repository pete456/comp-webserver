class User {
	constructor(user,pass,title) {
		this.user=user
		this.pass=pass
		this.title=title
	}
}

function SendCreds() {
	console.log("sending creds")
	var user = new User(document.getElementById("unametxt").value, document.getElementById("passtxt").value, document.getElementById("titletxt").value)

	var htmlreq = new XMLHttpRequest();
	htmlreq.open("POST","/adduser")
	htmlreq.onreadystatechange = function() {
		if(htmlreq.readyState === XMLHttpRequest.DONE){
			window.location="/login.html"	
		}
	}
	htmlreq.setRequestHeader("Content-type","application/json")
	var jsonstr = JSON.stringify(user)
	console.log(jsonstr)
	htmlreq.send(jsonstr)
}

var createbtn = document.getElementById("createbtn")
createbtn.addEventListener("click", SendCreds)
