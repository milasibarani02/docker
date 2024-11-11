menjalankan postgres sbg container dngn nama container "my-postgres-milasibarani02" dngn env dan volume "my-pg-volume-milasibarani02"
![image](https://github.com/user-attachments/assets/d23ff3b2-7bae-4190-8fbb-8687eaaef183)

Cek postgres sudah bisa diakses dbeaver
![image](https://github.com/user-attachments/assets/5d8e0cb1-d766-429f-b0bb-4b8ed159bb21)

Membuat table baru via dbeaver dngn nama “new_table_milasibarani“, dan menambah record
![image](https://github.com/user-attachments/assets/e1972a09-503d-48cf-9736-0e01f6d52e44)

Stop dan Hapus container postgres "my-postgres-milasibarani02“, 
dan menjalankan ulang postgres sbg container dengan nama container yg BARU "my-postgres-v2-milasibarani" dengan env dan volume yg sama
![image](https://github.com/user-attachments/assets/ae0ccba9-d702-460c-96bd-5a8e9bc486f9)

Cek table dan data masih ada
![image](https://github.com/user-attachments/assets/63511336-3d54-41fa-83b9-d911146690c7)
