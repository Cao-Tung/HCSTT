<?php

error_reporting(E_ALL);

require_once __DIR__.'/php/lib/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;

$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift', __DIR__ . '/php/lib');
$loader->register();

use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\THttpClient;
use Thrift\Transport\TBufferedTransport;
use Thrift\Exception\TException;
use gen_php\Hello;
include_once 'gen_php/Hello.php';
try {
  $socket = new TSocket('127.0.0.1', 1996);
  $transport = new TBufferedTransport($socket, 1024, 1024);
  $protocol = new TBinaryProtocol($transport);
  $client = new HelloClient($protocol);
  $transport->open();

  print $client->HelloString("Hello PHP");
  print "\n";

} catch (TException $tx) {
  print 'TException: '.$tx->getMessage()."\n";
}
?>