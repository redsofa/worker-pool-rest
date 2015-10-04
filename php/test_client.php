<?php
require 'vendor/autoload.php';
use GuzzleHttp\Client;

$json = '{"table": [{ "Index": 0, "NumA": 1, "NumB": 2 },{ "Index": 1, "NumA": 3,"NumB": 4}]}';

$client = new Client();
$url = 'http://localhost:8080/calcs';

$response = $client->post($url, ['body' => $json]);
echo ($response->getBody());
?>
