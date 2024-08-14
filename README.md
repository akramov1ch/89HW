# Uy vazifasi: Go yordamida Kafka bilan Mikroservis yaratish

## Maqsad
Ushbu uyga vazifada siz `Kafka` bilan ishlaydigan kichik mikroservis arxitekturasini yaratishingiz kerak bo'ladi. Asosiy maqsad `Kafka`, `Go` va mikroservislar bilan amaliy tajriba orttirishdir.


## Talablar
1. **Vazifa 1: Kafka Topic Yaratish**:
    -  3 ta partition va 1 ta replication olish faktori bilan `order_events` nomli topic ni yarating.

2. **Vazifa 2: Producer Servisini amalga oshirish:**:
    - `Kafka` producer sifatida ishlaydigan `Go` asosidagi mikroservis (`order-service`) ishlab chiqing.
    - Ushbu servis buyurtma ma'lumotlarini qabul qiluvchi, qayta ishlovchi va `order_events` `Kafka` topic ga xabar yuboruvchi `REST API` endpointini taqdim etishi kerak.

3. **Vazifa 3: Consumer Servisini amalga oshirish:**:
    - `Kafka` consumer sifatida ishlaydigan yana bir `Go` asosidagi mikroservis (`notification-service`) ishlab chiqing.
    - Ushbu servis `order_events` topic dan xabarlarni iste'mol qilishi, buyurtmalarni qayta ishlashi va harakatni amalga oshirishi (masalan, bildirishnoma yuborishni simulyatsiya qilishi) kerak.

4. **Bonus Vazifa: Xabar Filtrlashni amalga oshiring**:
    - Consumer service da faqat ma'lum turdagi xabarlar (masalan, faqat "`new orders`") qayta ishlanishi uchun xabarlarni filtrlashni amalga oshiring.











