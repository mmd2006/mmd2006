#  ToDoApp

یک RESTful API ساده برای مدیریت تسک‌ها (Tasks) با استفاده از زبان Go، دیتابیس MongoDB و فریمورک Echo.  
این پروژه شامل احراز هویت JWT، مدیریت نقش (کاربر / ادمین)، و عملیات کامل CRUD می‌باشد.

---

##  تکنولوژی‌های استفاده‌شده

- [Go](https://golang.org/)
- [Echo Framework](https://echo.labstack.com/)
- [MongoDB](https://www.mongodb.com/)
- [mongo-go-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [Postman](https://www.postman.com/) – برای تست API

---

## ⚙️ نصب و اجرا

```bash
# 1. کلون کردن پروژه
git clone https://github.com/mmd2006/mmd2006
cd ToDoApp

# 2. اضافه کردن فایل .env با محتویات زیر:
JWT_SECRET=your_secret_key
MONGODB_URI=mongodb+srv://reza1385312:JC9T75d8oZGPbqaa@cluster0.j7sguxp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0

# 3. نصب پکیج‌ها
go mod tidy

# 4. اجرای پروژه
go run main.go

```
---

| متد    | مسیر           | توضیح                          | سطح دسترسی    |
| ------ | -------------- | ------------------------------ | ------------- |
| POST   | `/signup`      | ساخت حساب کاربری جدید          | عمومی         |
| POST   | `/login`       | دریافت توکن JWT برای ورود      | عمومی         |
| GET    | `/tasks`       | دریافت تمام تسک‌های کاربر فعلی | کاربر / ادمین |
| POST   | `/tasks`       | ساخت تسک جدید                  | کاربر / ادمین |
| GET    | `/tasks/:id`   | دریافت تسک خاص متعلق به کاربر  | کاربر / ادمین |
| PUT    | `/tasks/:id`   | ویرایش تسک (در صورت مالک بودن) | کاربر / ادمین |
| DELETE | `/tasks/:id`   | حذف تسک (در صورت مالک بودن)    | کاربر / ادمین |
| GET    | `/admin/tasks` | دریافت تمام تسک‌ها             | فقط ادمین     |
| GET    | `/admin/users` | دریافت لیست کاربران            | فقط ادمین     |

---

Mohammadreza https://github.com/mmd2006/mmd2006