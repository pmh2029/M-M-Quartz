package user

type UserModule struct {
	Repository UserRepositoryInterface
	Schema     UserSchema
}
