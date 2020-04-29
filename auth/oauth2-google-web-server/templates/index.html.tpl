<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Golang examples OAuth2 with Google</title>
        <meta name="author" content="">
        <meta name="description" content="">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="google-signin-client_id" content="{{ .ClientID }}.apps.googleusercontent.com">

        <link href="css/normalize.css" rel="stylesheet">
        <link href="css/style.css" rel="stylesheet">
    </head>

    <body>
        <div class="g-signin2" data-onsuccess="onSignIn"></div>
        <a href="/login">sign in</a>

        <script src="https://apis.google.com/js/platform.js" async defer></script>
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
        <script src="js/script.js"></script>
    </body>
</html>