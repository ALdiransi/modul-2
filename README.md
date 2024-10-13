# MODUL-2_-REVIEW-STRUKTUR-KONTROL
2311102328_Aldi Ransi

-tugas 2A

1.Program ini akan menukar tiga input string yang dimasukkan oleh user,pertama tama user di minta memasukan ketiga string lalu pada string pertama di simpan ke temp lalu string ke dua di pindakan ke tempat pertama dan string ke tiga di pindakan ke temapt ke dua
lalu string yang di simpan di tempat sementara di pindah kan ke tempat ke 3

2.Program bertujuan menunjukkan bahwa suatu tahun inputan user termasuk dalam tahun kabisat atau bukan, dengan cara kerja pada bagian fungsi program akan membagikan inputan user menjadi empat ketika hasilnya bilangan real 
program akan menampilkan fmt.Println(tahun, "adalah tahun kabisat.") sedangkan jika tidak program akan mencetak fmt.Println(tahun, "bukan tahun kabisat.")

3.Program akan  menghitung volume dan luas permukaan sebuah bola berdasarkan input jari jari (radius) yang diberikan oleh user,dengan cara kerja program akan meminta inputan user jari jari dari bola dan dengan mengunakan fungsi math.Pow() untuk menghitung jari jari dan menggunakan math.Pi
untuk menghitung angka kostannya

4.Program mengubah suhu celsius inputan dari user menjadi reamur, kelvin, dan fahreinheit, dengan perintah ( reamur := celsius * 4 / 5)(kelvin := celsius + 273.15)(fahrenheit := (celsius * 9 / 5) + 32 
)

5.Program ini berfungsi untuk mengubah 2 baris inputan user, 1 baris pertama berisi 5 data integer yang akan dicetak dalam format karakter, baris kedua berisi 3 data char yang akan dicetak +1 setelah masing-masing char tersebut

-2B

1.Program ini memeriksa 4 urutan warna inputan user sudah sesuai dengan yang sudah di tentukan pada kodingan , jika sudah akan muncul boolean true dan sebaliknya.

2.Program ini berfungsi untuk mencetak n sebanyak  yang diinputkan user, nantinya user diminta untuk memasukkan nama bunga sebanyak n juga, yang kemudian akan dirangkai dengan "-" sesuai yang sudah diinputkan.

3.Program bertujuan mengecek apakah beban belanjaan yang dibawa oleh Pak Andi akan membuat motor oleng atau tidak dengan cara menentukan total beban minimal dan selisih dari beban kanan dan kiri, beban tersebut akan di input oleh user dan di periksa pada if totalBeban >= 100 && selisihBeban <= 10

4.Program ini berfungsi untuk nilai akar 2 dari suatu bilangan inputan user dengan Menggunakan fungsi math.Sqrt() untuk menghitung akar kuadrat dari bilangan tersebut.

-2C

1.Program ini sebagai menghitung biaya pengiriman Pertama-tama program meminta user untuk menginputkan berat parcel dalam gram tersebut. Setelah itu berat parsel dikonversi ke kilogram dan gram sisa dengan cara membagi berat total dengan 1000, biaya dasar dihitung berdasarkan berat kilogram, di mana setiap kilogram dikenakan biaya Rp. 10.000.

2.Program ini berfungsi untuk menghitung sebuah nilai akhir mata kuliah berdasarkan inputan dari sebuah user, dimana program ini akan menampilkan nilai akhir berdasarkan abjad "A,AB,B,BC,C,C,E" Soal & Jawaban.

i.Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal? Jawaban: Program eror, dikarenakan program tersebut tidak konsisten, terutama pada variablel nam dideklarasikan sebagai float64, namun di dalam blok if, program mencoba mengassign string (seperti "A", "AB", dll.) ke nam.

ii.Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya! Jawaban:
Kesalahan program:

-Program tidak konsisten, pada variablel nam dideklarasikan sebagai float64, namun di dalam blok if, program mencoba memasukkan string (seperti "A", "AB", dll.) ke nam.
-Kondisi if adalah independen, sehingga jika nam lebih besar dari beberapa batas, beberapa blok if akan dieksekusi secara berurutan, dan nilai akhir nmk akan ditimpa beberapa kali.
-Sebagai contoh, jika nam = 80.1, maka semua kondisi berikut akan terpenuhi: nam > 80 → nmk = "A" ,nam > 72.5 → nmk = "AB" ,nam > 65 → nmk = "B" ,nam > 57.5 → nmk = "BC" ,nam > 50 → nmk = "C" ,nam > 40 → nmk = "D" Akhirnya, nilai nmk akan menjadi "D", yang jelas tidak sesuai dengan spesifikasi yang diharapkan.

Perbaikan Program:

*Gantilah nam = "A" dengan nmk = "A", dan seterusnya untuk semua kondisi.

*Gunakan else if untuk memastikan hanya satu kondisi yang terpenuhi dan mencegah nilai nmk ditimpa oleh kondisi yang berikutnya.

iii.Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'. Jawaban: Input: 93.5 -> Output: A. Input: 70.6 -> Output: B. Input: 49.5 -> Output: D.

3.Program ini sebagai membuktikkan suatu bilangan inputan dari user termasuk kedalam bilangan prima atau bukan, dengan cara mencari faktor-faktor nya.
