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
MONGODB_URI=mongodb+srv://reza1385312:JC9T75d8oZGPbqaa@cluster0.j7sguxp.mongodb.net/?appName=Cluster0
PORT=1323

# 3. نصب پکیج‌ها
go mod tidy

# 4. اجرای پروژه
go run main.go

```
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
```
---

# 🔑 حساب کاربری پیش‌فرض ادمین

برای ورود با ادمین:

Username: admin

Password: admin123

URL: http://localhost:1323

پس از لاگین، توکن JWT دریافت می‌کنید که برای تمام درخواست‌های محافظت‌شده لازم است.

---

# 📦 ساختار کالکشن‌ها (MongoDB)

## Task Collection (tasks)
| فیلد          | نوع      | توضیح                |
| ------------- | -------- | -------------------- |
| `_id`         | ObjectID | شناسه یکتا           |
| `user_id`     | ObjectID | شناسه کاربر مالک تسک |
| `title`       | string   | عنوان تسک            |
| `description` | string   | توضیحات (اختیاری)    |
| `completed`   | bool     | وضعیت انجام شدن تسک  |

---

## User Collection (users)
| فیلد       | نوع      | توضیح                    |
| ---------- | -------- | ------------------------ |
| `_id`      | ObjectID | شناسه یکتا               |
| `username` | string   | نام کاربری               |
| `password` | string   | هش رمز عبور              |
| `role`     | string   | نقش کاربر (`user/admin`) |


---

## مسیرهای API
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