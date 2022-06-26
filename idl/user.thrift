namespace go user

struct LoginRequest {
    1: string Username;
    2: string Password;
}

struct LoginResponse {
    1: string Message;
}

service Login {
    LoginResponse CheckUser(1: LoginRequest req);
}