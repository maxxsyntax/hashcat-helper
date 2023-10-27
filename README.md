# hashcat-masks
Masks for Hashcat mode 22000 and 22001.  
Regional CPE was distributed with a default keyspace that can be reduced to 256. 
A OUI of the vendor and the last 2 bytes of the BSSID were used as a consistent string for key generation.  This leaves only 1 byte (256) to be recoverd. 
