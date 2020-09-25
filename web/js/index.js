function formatTitlesAndUsers() {
	let values = JSON.parse(this.response)
	let pl = document.getElementById("portal-list")
	for(e of values.PD) {
		let li = document.createElement("li")
		let a = document.createElement("a")
		a.href = "site/"+e.Name
		a.innerHTML = e.Title
		li.appendChild(a)
		pl.appendChild(li)
		console.log(e)
	}
}

function getTitlesAndUsers() {
	var req = new XMLHttpRequest()
	req.addEventListener("load", formatTitlesAndUsers)
	req.open("GET","/getportaldata")
	req.send()
}

getTitlesAndUsers()
