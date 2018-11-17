# googlespeech2text
google speech to text sample

Run SOX in mac to convert to Linear16  

//Raw local audio file in "wav" format  
mbp-008814:GCP_AI sk$ ls -lrt sk1a.wav  
-rw-r--r--@ 1 sk  AD\Domain Users  1025792 Nov 16 14:17 sk1a.wav  
mbp-008814:GCP_AI sk$

//Convert to LINEAR16  
mbp-008814:GCP_AI sk$ sox sk1a.wav --channels=1 --rate 16k --bits 16 sk1a.raw

mbp-008814:GCP_AI sk$ ls -lrt sk1a.raw   
-rw-r--r--  1 sk  Users  1021696 Nov 16 14:17 sk1a.raw
mbp-008814:GCP_AI sk$ 

//Result from the speech below: Sample code attached  
mbp-008814:GCP_AI sk$ ./speechtotext -input sk1a.raw  
"four score and seven years ago our fathers brought forth on this continent a new nation conceived in liberty and dedicated to the proposition that all men are created equal now we are engaged in a great civil war testing whether that Nation or any Nation so conceived and so dedicated can long endure we are met on a great battlefield of that war we have come to dedicate a portion of that field as a final resting place for those who here gave their lives that that Nation might live it is also it is altogether fitting and proper that we should do this" (confidence=0.979804)  
mbp-008814:GCP_AI sk$
