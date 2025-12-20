# ToDoApp

یک RESTful API ساده برای مدیریت تسک‌ها (Tasks) با زبان Go، دیتابیس MongoDB و فریمورک Echo.  
این پروژه شامل احراز هویت JWT، مدیریت نقش (کاربر / ادمین) و عملیات کامل CRUD است.

---

## تکنولوژی‌های استفاده‌شده

- [Go](https://golang.org/)
- [Echo Framework](https://echo.labstack.com/)
- [MongoDB](https://www.mongodb.com/)
- [mongo-go-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [Postman](https://www.postman.com/) – برای تست API

---

## نصب و اجرا

```bash
# 1. کلون کردن پروژه
git clone https://github.com/mmd2006/mmd2006
cd ToDoApp

# 2. اضافه کردن فایل .env (نمونه)
# ⚠️ لطفاً اطلاعات حساس واقعی را وارد نکنید، فقط نمونه:
JWT_SECRET=your_secret_key
MONGODB_URI=your_mongodb_uri_here
PORT=1323

# 3. نصب پکیج‌ها
go mod tidy

# 4. اجرای پروژه
go run main.go

ساختار پروژه
go
```plaintext
📁 ToDoApp
├─ main.go
├─ go.mod
├─ .env
├─ config/
│  └─ mongo.go
├─ controller/
│  ├─ task.go
│  └─ user.go
├─ middleware/
│  └─ jwt.go
├─ model/
│  ├─ task.go
│  └─ user.go
├─ router/
│  └─ router.go
└─ validation/
   ├─ task.go
   └─ user.go
...   

| متد    | مسیر           | توضیح                          | سطح دسترسی    |
| ------ | -------------- | ------------------------------ | ------------- |
| POST   | `/signup`      | ساخت حساب کاربری جدید          | عمومی         |
| POST   | `/login`       | دریافت توکن JWT برای ورود      | عمومی         |
| GET    | `/tasks`       | دریافت تمام تسک‌های کاربر فعلی | کاربر / ادمین |
| POST   | `/tasks`       | ساخت تسک جدید                  | کاربر / ادمین |
| GET    | `/tasks/:id`   | دریافت تسک خاص متعلق به کاربر  | کاربر / ادمین |
| PUT    | `/tasks/:id`   | ویرایش تسک (در صورت مالک بودن)| کاربر / ادمین |
| DELETE | `/tasks/:id`   | حذف تسک (در صورت مالک بودن)   | کاربر / ادمین |
| GET    | `/admin/tasks` | دریافت تمام تسک‌ها             | فقط ادمین     |
| GET    | `/admin/users` | دریافت لیست کاربران            | فقط ادمین     |


توضیح: مسیرهای /admin/... فقط برای ادمین و مسیرهای /tasks/... برای کاربر لاگین شده قابل دسترسی هستند.

تست API با Postman
پروژه را اجرا کنید (go run main.go)

از Postman برای تست endpoint ها استفاده کنید:

ابتدا /signup یا /login برای دریافت JWT

سپس JWT را در header به شکل Authorization: Bearer <token> قرار دهید

بعد از آن می‌توانید مسیرهای /tasks و /admin/... را تست کنید

امنیت
JWT برای احراز هویت و Role-based Access Control برای مدیریت نقش‌ها

رمز عبور کاربران با bcrypt هش شده و در دیتابیس ذخیره می‌شود

فایل .env شامل اطلاعات حساس است و نباید در گیت‌هاب قرار گیرد

لینک‌ها
GitHub: https://github.com/mmd2006/mmd2006