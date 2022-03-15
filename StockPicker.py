x = input("banyak stok: ")
y = []

for i in range(int(x)):
	z = input("nilai stok: ")
	y.append(z)

kecil = int(y[0])
pengeluaran = 0
maksimal = 0

for j in y:
	kecil = min(kecil, int(j))
	pengeluaran = int(j) - kecil
	maksimal = max(maksimal, pengeluaran)

print(maksimal)
