puts "crypto.js"
`java -jar ../compiler.jar --js=aes.js --js=md5.js --js=sha1.js --js=pad-zeropadding.js --js=enc-base64.js --js=rsa/jsbn.js --js=rsa/prng4.js --js=rsa/rng.js --js=rsa/rsa.js --js_output_file=../../crypto.js`

#puts "rsa.js"
#`java -jar ../compiler.jar --js=rsa/jsbn.js --js=rsa/prng4.js --js=rsa/rng.js --js=rsa/rsa.js --js_output_file=../rsa.js`
