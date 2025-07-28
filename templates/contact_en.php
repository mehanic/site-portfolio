<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Contact Form</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #191970;
            margin: 0;
            padding: 0;
        }

        .container {
            width: 50%;
            margin: 30px auto;
            padding: 20px;
            background: white;
            border-radius: 8px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            text-align: center;
        }

        h2 {
            color: #2c3e50;
        }

        form {
            display: flex;
            flex-direction: column;
        }

        label {
            font-weight: bold;
            margin-top: 10px;
            text-align: left;
        }

        input, textarea {
            width: 100%;
            padding: 10px;
            margin-top: 5px;
            border: 1px solid #ccc;
            border-radius: 5px;
            box-sizing: border-box;

        }

        button {
            background-color: #3498db;
            color: white;
            border: none;
            padding: 10px;
            margin-top: 15px;
            border-radius: 5px;
            cursor: pointer;
        }

        button:hover {
            background-color: #2980b9;
        }
        p,a{color:white;}
    </style>
</head>
<body>
<header style="display: flex; justify-content: center;">
<IMG src="bg/1.jpg" id="head_img" onclick="doRotate();">
</header>
<SCRIPT>
image=[];
for(i=1;i<=10;i++){
	image[i] = new Image();
	image[i].src="bg/"+i+".jpg";
}
current=1;
function doRotate(){
	current = current%10+1;
	document.getElementById('head_img').src=image[current].src;
}
</SCRIPT>

    <div class="container">
        
<?php
$name=$_REQUEST['name'];
$email=$_REQUEST['email'];
$message=$_REQUEST['message'];
if($message){
	$ip=$_SERVER['HTTP_X_FORWARDED_FOR']?$_SERVER['HTTP_X_FORWARDED_FOR']:$_SERVER['REMOTE_ADDR'];
	include("SxGeo.php");
	$SxGeo = new SxGeo('SxGeoCity.dat');
	$country = $SxGeo->getCountry($ip);
	$geo = $SxGeo->getCityFull($ip);
	unset($SxGeo);

	$info="\n\n".$message."\n\n---\n";
	$info.="\nName       ".$name;
	$info.="\nEmail      ".$email;
	$info.="\nCountry    ".$geo['country']['name_en']." (".$geo['country']['iso'].")";
	$info.="\nCity       ".$geo['city']['name_en']." [".$geo['city']['lat'].",".$geo['city']['lon']."]";
	$info.="\n\nIP         ".$ip;
	$info.="\nPROVIDER   ".gethostbyaddr($ip);
	$info.="\nPROXY NAME ".($_SERVER['HTTP_VIA']?$_SERVER['HTTP_VIA']:$_SERVER['HTTP_FORWARDED']);
	$info.="\nPROXY IP   ".($_SERVER['HTTP_X_FORWARDED_FOR']?$_SERVER['REMOTE_ADDR']:'');
	$info.="\nSOFTWARE   ".$_SERVER['HTTP_USER_AGENT'];
	
	$headers="From: ".$email." \nContent-Type: text/plain; charset=utf-8\nContent-Transfer-Encoding: 8bit";
	mail("mehanic2000@gmail.com", "CV Request", $info, $headers);
	echo "<textarea rows='20' style='width:100%'>".$info."</textarea>";

	echo "<br><br><h2>Thank you, your message has been sent</h2><br><br>";
}else{
?>
        
        <h2>Send a Message</h2>
        <form action="<?=$_SERVER['PHP_SELF']?>" method="POST">
            <label for="name">Name *</label>
            <input type="text" id="name" name="name" required>

            <label for="email">Email *</label>
            <input type="email" id="email" name="email" required>

            <label for="message">Message *</label>
            <textarea id="message" name="message" rows="5" required></textarea>

            <button type="submit">Send Message</button>
        </form>
<?}?>
    </div>
    <footer style="display: flex; justify-content: center;">
        <p>&copy;2025<BR> ✉ <a href="mailto:mehanic2000@gmail.com">mehanic2000@gmail.com</a></p>
    </footer>

</body>
</html>