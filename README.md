# Kripto Veri API'si

Bu uygulama, kripto para fiyat verilerini sağlamak için kullanılan bir API sunucusudur. Binance borsasından canlı fiyat verilerini alır ve istemcilere sunar.

## Kurulum

1. Bu kodu çalıştırmak için öncelikle Go programlama dilinin yüklü olması gerekmektedir.
2. MongoDB veritabanı sunucusunun yüklü ve çalışır durumda olması gerekmektedir.
3. Terminal veya komut istemcisinde projenin bulunduğu dizine gidin.
4. `go run main.go` komutunu kullanarak uygulamayı başlatın.
5. Tarayıcınızdan `http://localhost:8080/price` adresine giderek kripto para fiyatlarını görüntüleyin.

## Kullanım

- `/price` endpoint'i üzerinden kripto para fiyatlarını alabilirsiniz. Veriler önbelleklenir ve her istek arasında son alınan veri kullanılır. Veri, MongoDB'den alınır ve her istekte önbelleğe alınır.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.
