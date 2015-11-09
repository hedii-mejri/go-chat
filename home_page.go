package main

import(
  "html/template"
  "net/http"
  "fmt"
  "encoding/json"
)

const login_page = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8"/>
  <title> Go-Chat </title>
  <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
  <link rel="stylesheet" href="/assets/stylesheets/login.css">
  <link href='http://fonts.googleapis.com/css?family=Montserrat:400,700' rel='stylesheet' type='text/css'>
  <script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
  <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
  <script src="/assets/javascripts/login.js"></script>
</head>
<body style="background-color: #FAFAFF;">
  <div class="login-block">
    <h1>Login</h1>
    <input type="text" value="" placeholder="Username" id="username" />
    <button id="login_button">Submit</button>
  </div>
</body>
</html>
`

const home_page = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8"/>
  <title> Go-Chat </title>
  <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
  <link rel="stylesheet" href="/assets/stylesheets/emojify.min.css" />
  <link rel="stylesheet" href="/assets/stylesheets/magnific-popup.min.css">
  <link rel="stylesheet" type="text/css" href="/assets/stylesheets/chat.css">
  <script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
  <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
  <script src="/assets/javascripts/emojify.min.js"></script>
  <script src="/assets/javascripts/jquery.ui.widget.min.js"></script>
  <script src="/assets/javascripts/jquery.iframe-transport.min.js"></script>
  <script src="/assets/javascripts/jquery.fileupload.min.js"></script>
  <script src="/assets/javascripts/jquery.magnific-popup.min.js"></script>
  <script src="/assets/javascripts/chat.js"></script>
</head>
<body style="background-color: #FAFAFF;">
  <h1> Just Go-Chat</h1>
  <div id="data" style="display: none;">{{.}}</div>
  <div class="container-fluid">
    <div class="panel panel-primary">
      <div class="panel-heading">
        <span class="glyphicon glyphicon-comment"></span> Chat Box
        <div class="btn-group pull-right">
          <button type="button" class="btn btn-default btn-xs dropdown-toggle" data-toggle="dropdown">Options
            <span class="glyphicon glyphicon-chevron-down"></span>
          </button>
          <ul class="dropdown-menu slidedown">
            <li id="clear_chat"><a href="#"><span class="glyphicon glyphicon-unchecked">
            </span>Clear Chat</a></li>
            <li id="leave_chat"><a href="#"><span class="glyphicon glyphicon-remove-sign">
            </span>Leave Chat</a></li>
            <li id="join_chat"><a href=""><span class="glyphicon glyphicon-ok-sign">
            </span>Join Chat</a></li>
          </ul>
        </div>
      </div>
      <div id ="chat_box" class="panel-body msg_container_base"></div>

      <div class="panel-footer">
        <div class="row">
          <div class="col-lg-9">
            <input id="chat_prompt" type="text" class="form-control input-sm" placeholder="Type your message here..." />
          </div>
          <div class="col-lg-3 select-wrapper">
            <input id="fileupload" type="file" title="add files" name="files" data-url="/upload" multiple accept="*"/>
          </div>

          <div class="col-lg-3">
              <button class="btn btn-warning btn-sm" id="send" style="display: none;">Send</button>
          </div>

          <div class="progress">
            <div class="progress-bar progress-bar-striped active" role="progressbar" aria-valuemin="0" aria-valuemax="100" style="width:0%"/>
          </div>
        </div>
      </div>
    </div>
  </div>
  <audio id="notify" preload="auto">
    <source src="http://demos.9lessons.info/notify/notify.ogg" type="audio/ogg">
    <source src="http://demos.9lessons.info/notify/notify.mp3" type="audio/mpeg">
    <source src="http://demos.9lessons.info/notify/notify.wav" type="audio/wav">
  </audio>

</body>
</html>
`

func home_handler(w http.ResponseWriter, r *http.Request){
  name := r.FormValue("name")
  if name == "" {
    name = "anonymous person"
  }
  host := fmt.Sprintf("%s:%d", *hostname, *port)
  parsedUrl := map[string]string{"context": host, "name": name }
  json, _ := json.Marshal(parsedUrl)
  home_template, error := template.New("webpage").Parse(home_page)
  if error != nil {
    fmt.Println("Couldn't parse home page!")
  }
  err := home_template.Execute(w, string(json))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func login_handler(w http.ResponseWriter, r *http.Request){
  home_template, error := template.New("webpage").Parse(login_page)
  if error != nil {
    fmt.Println("Couldn't parse home page!")
  }
  err := home_template.Execute(w, nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func redirect_handler(w http.ResponseWriter, r *http.Request){
  http.Redirect(w, r, "/login", http.StatusFound)
}
