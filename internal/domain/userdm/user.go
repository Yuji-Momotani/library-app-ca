// package userdm

// import (
// 	"errors"
// 	"time"
// )

// type UserStatus string

// const (
// 	UserStatusActive    UserStatus = "active"
// 	UserStatusSuspended UserStatus = "suspended"
// )

// // Userエンティティ - 図書館の利用者を表す
// type User struct {
// 	id          UserID
// 	name        string
// 	email       string
// 	status      UserStatus
// 	overdueFees float64
// 	createdAt   time.Time
// 	// No currentLoanCount field - derive it from the Loan table
// }

// const MaxLoans = 5

// // NewUser creates a new active user
// func NewUser(name, email string) *User {
// 	return &User{
// 		id:          NewUserID(),
// 		name:        name,
// 		email:       email,
// 		status:      UserStatusActive,
// 		overdueFees: 0,
// 		createdAt:   time.Now(),
// 	}
// }

// // ReconstructUser rebuilds user from persistence
// // // Q：RecnstructXXX はRepositoryでの再構築専用の関数だと思うが、
// // // アプリケーションレイヤーから直接使用されると、ガード節が効かずに
// // // 予期せぬ値がセットされる可能性が考えられるが、そこはどう思うか？
// // // DBのデータにも不正値があるかもよ？
// func ReconstructUser(
// 	id UserID,
// 	name, email string,
// 	status UserStatus,
// 	overdueFees float64,
// 	createdAt time.Time,
// ) *User {
// 	return &User{id, name, email, status, overdueFees, createdAt}
// }

// // User entity focuses on RULES, not STATE
// // Loan count is derived from the Loan table

// // // Q：ドメインサービスにおくべきでは？
// // // currentLoanCountはLoanドメインに属する概念？→その場合、ドメインをまたぐチェックのため、
// // // ドメインサービスが良いのではないか？
// func (u *User) CanBorrow(currentLoanCount uint32) bool {
// 	if u.status == UserStatusSuspended {
// 		return false
// 	}
// 	if currentLoanCount >= MaxLoans {
// 		return false
// 	}
// 	if u.overdueFees > 0 {
// 		return false
// 	}
// 	return true
// }

// // State changes are tracked in Loan table (no BorrowBook/ReturnBook methods needed)
// // // Q：なぜ新たなUserを返している？
// // // u.overdueFees += amount だけで良くないか？（レシーバーがポインタ型だし、entityはライフサイクルのあるオブジェクトなので）
// func (u *User) AddOverdueFee(amount float64) (*User, error) {
// 	if amount <= 0 {
// 		return nil, errors.New("overdue fee must be greater than 0")
// 	}
// 	return &User{
// 		id:          u.id,
// 		name:        u.name,
// 		email:       u.email,
// 		status:      u.status,
// 		overdueFees: u.overdueFees + amount,
// 		createdAt:   u.createdAt,
// 	}, nil
// }

// func (u *User) PayOverdueFee(amount float64) (*User, error) {
// 	if amount < 0 {
// 		return nil, errors.New("payment amount cannot be negative")
// 	}
// 	if amount > u.overdueFees {
// 		return nil, errors.New("payment exceeds current overdue fees")
// 	}
// 	return &User{
// 		id:          u.id,
// 		name:        u.name,
// 		email:       u.email,
// 		status:      u.status,
// 		overdueFees: u.overdueFees - amount,
// 		createdAt:   u.createdAt,
// 	}, nil
// }

// // ゲッター（Goの慣習: "Get"プレフィックスなし）
// func (u *User) Id() UserID {
// 	return u.id
// }

// func (u *User) Name() string {
// 	return u.name
// }

// func (u *User) Email() string {
// 	return u.email
// }

// func (u *User) Status() UserStatus {
// 	return u.status
// }

// func (u *User) OverdueFees() float64 {
// 	return u.overdueFees
// }

// func (u *User) CreatedAt() time.Time {
// 	return u.createdAt
// }

// // Note: CurrentLoanCount is NOT stored in the User entity
// // Use LoanRepository.CountActiveLoansForUser() to get the current count

// // HasOverdueBooks checks if user has overdue books (overdue fees > 0)
// func (u *User) HasOverdueBooks() bool {
// 	return u.overdueFees > 0
// }

package userdm

import (
	"errors"
	"time"
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
)

type User struct {
	id          *UserID
	name        string
	email       string
	status      UserStatus
	overdueFees float64
	createdAt   time.Time
}

const MaxLoans = 5

// NewUser は新しいアクティブなユーザーを作成します
func NewUser(name, email string) *User {
	return &User{
		id:          GenerateUserID(),
		name:        name,
		email:       email,
		status:      UserStatusActive,
		overdueFees: 0,
		createdAt:   time.Now(),
	}
}

// ReconstructUser は永続化からユーザーを再構築します
func ReconstructUser(
	id *UserID,
	name, email string,
	status UserStatus,
	overdueFees float64,
	createdAt time.Time,
) *User {
	return &User{id, name, email, status, overdueFees, createdAt}
}

// 提供された貸出数に基づいて貸出可能性を検証
// 貸出数はユースケースによって提供される（Loanテーブルから導出）
func (u *User) CanBorrow(currentLoanCount uint32) bool {
	if u.status == UserStatusSuspended {
		return false
	}
	if currentLoanCount >= MaxLoans {
		return false
	}
	if u.overdueFees > 0 {
		return false
	}
	return true
}

// 延滞料金の追加
func (u *User) AddOverdueFee(amount float64) (*User, error) {
	if amount <= 0 {
		return nil, errors.New("延滞料金は0より大きい必要があります")
	}
	return &User{
		id:          u.id,
		name:        u.name,
		email:       u.email,
		status:      u.status,
		overdueFees: u.overdueFees + amount,
		createdAt:   u.createdAt,
	}, nil
}

// 不変の状態変更: 延滞料金を支払う
func (u *User) PayOverdueFee(amount float64) (*User, error) {
	if amount < 0 {
		return nil, errors.New("支払い金額は負の値にできません")
	}
	if amount > u.overdueFees {
		return nil, errors.New("支払い金額が現在の延滞料金を超えています")
	}
	return &User{
		id:          u.id,
		name:        u.name,
		email:       u.email,
		status:      u.status,
		overdueFees: u.overdueFees - amount,
		createdAt:   u.createdAt,
	}, nil
}

func (u *User) Suspend() *User {
	return &User{
		id:          u.id,
		name:        u.name,
		email:       u.email,
		status:      UserStatusSuspended,
		overdueFees: u.overdueFees,
		createdAt:   u.createdAt,
	}
}

// 不変の状態変更: アカウントを有効化
func (u *User) Activate() *User {
	return &User{
		id:          u.id,
		name:        u.name,
		email:       u.email,
		status:      UserStatusActive,
		overdueFees: u.overdueFees,
		createdAt:   u.createdAt,
	}
}

// ゲッター
func (u *User) Id() *UserID          { return u.id }
func (u *User) Name() string         { return u.name }
func (u *User) Email() string        { return u.email }
func (u *User) Status() UserStatus   { return u.status }
func (u *User) OverdueFees() float64 { return u.overdueFees }
func (u *User) CreatedAt() time.Time { return u.createdAt }
