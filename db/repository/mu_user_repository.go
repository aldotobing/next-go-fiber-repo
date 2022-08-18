package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/str"
)

type IMuUserRepository interface {
	FindAll(c context.Context, parameter models.MuUserParameter) ([]models.MuUser, int, error)
	FindByID(c context.Context, parameter models.MuUserParameter) (models.MuUser, error)
	FindByRefferalCode(c context.Context, parameter models.MuUserParameter) (models.MuUser, error)
	FindByEmail(c context.Context, parameter models.MuUserParameter) (models.MuUser, error)
	FindByQrCode(c context.Context, parameter models.MuUserParameter) (models.MuUser, error)
	SetActiveUser(c context.Context, model *models.MuUser) (string, error)
	Add(c context.Context, model *models.MuUser) (*string, error)
	Edit(c context.Context, model *models.MuUser) (string, error)
	Delete(c context.Context, id string, user_id string, now time.Time) (string, error)
	UpdatePassword(c context.Context, model *models.MuUser) (string, error)

	CheckRefferalCodeMaxLimit(c context.Context, parameter models.MuUserParameter) (bool, error)
}

type MuUserRepository struct {
	DB *sql.DB
}

func NewMuUserRepository(DB *sql.DB) IMuUserRepository {
	return &MuUserRepository{DB: DB}
}

// Scan rows
func (repository MuUserRepository) scanRows(rows *sql.Rows) (res models.MuUser, err error) {
	err = rows.Scan(
		&res.ID, &res.BranchID, &res.FbId, &res.GoogleId, &res.AppleId, &res.Name,
		&res.UserName, &res.Password, &res.Gender, &res.QrCode, &res.Level, &res.ReferalCode,
		&res.Email, &res.NoTelp, &res.Address, &res.RoleGroupId,
		&res.BirthDate, &res.UserName, &res.BranchCoverageStr,
		&res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.CreatedBy,
		&res.UpdatedBy, &res.DeletedBy, &res.UserName, &res.ReferralCodeLimitUse, &res.ImgKTP, &res.ImgProfile,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Scan row
func (repository MuUserRepository) scanRow(row *sql.Row) (res models.MuUser, err error) {
	err = row.Scan(
		&res.ID, &res.BranchID, &res.FbId, &res.GoogleId, &res.AppleId, &res.Name,
		&res.UserName, &res.Password, &res.Gender, &res.QrCode, &res.Level, &res.ReferalCode,
		&res.Email, &res.NoTelp, &res.Address, &res.RoleGroupId,
		&res.BirthDate, &res.UserName, &res.BranchCoverageStr,
		&res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.CreatedBy,
		&res.UpdatedBy, &res.DeletedBy, &res.UserName, &res.ReferralCodeLimitUse, &res.ImgKTP, &res.ImgProfile,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository MuUserRepository) FindByRefferalCode(c context.Context, parameter models.MuUserParameter) (data models.MuUser, err error) {
	statement := models.MuUserSelectStatement + ` WHERE def.deleted_at_user IS NULL AND def.referral_code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.ReferalCode)
	fmt.Println(statement)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository MuUserRepository) CheckRefferalCodeMaxLimit(c context.Context, parameter models.MuUserParameter) (allowed bool, err error) {

	statement := `select (
		case when 
		(select count(id_user_givereferral) from mt_referralcode where id_user_givereferral= $1)
		 <
		(select referral_limit_use from mu_user where id_user = $1)
		then true else false end
	   ) as available`

	err = repository.DB.QueryRowContext(c, statement, parameter.ID).Scan(&allowed)

	if err != nil {
		return allowed, err
	}

	return allowed, nil
}

func (repository MuUserRepository) FindByEmail(c context.Context, parameter models.MuUserParameter) (data models.MuUser, err error) {
	statement := models.MuUserSelectStatement + ` WHERE  def.email_user = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Email)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository MuUserRepository) FindByQrCode(c context.Context, parameter models.MuUserParameter) (data models.MuUser, err error) {
	statement := models.MuUserSelectStatement + ` WHERE def.deleted_at_user IS NULL AND def.qr_code = $1`
	row := repository.DB.QueryRowContext(c, statement, parameter.Email)

	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository MuUserRepository) FindByID(c context.Context, parameter models.MuUserParameter) (data models.MuUser, err error) {
	statement := models.MuUserSelectStatement + ` WHERE def.deleted_at_user IS NULL AND def.id_user = $1`

	row := repository.DB.QueryRowContext(c, statement, parameter.ID)
	data, err = repository.scanRow(row)
	if err != nil {

		return data, err
	}

	return data, nil
}

func (repository MuUserRepository) FindAll(c context.Context, parameter models.MuUserParameter) (data []models.MuUser, count int, err error) {
	conditionString := ``

	conditionString += ` AND def.role_group_id in (
		select distinct(role_group_id) from mp_role_group_line 
			where role_id not in (6,7,8,9) and active =1
	) `

	statement := models.MuUserSelectStatement + ` ` + models.MuUserWhereStatement +
		` AND (LOWER(def."name") LIKE $1 ) ` + conditionString + ` ORDER BY ` + parameter.By + ` ` + parameter.Sort + ` OFFSET $2 LIMIT $3`

	rows, err := repository.DB.QueryContext(c, statement, "%"+parameter.Search+"%", parameter.Offset, parameter.Limit)
	if err != nil {
		return data, count, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}

		data = append(data, temp)
	}

	statement = `SELECT COUNT(*) FROM "mu_user" def ` + models.MuUserWhereStatement + ` ` +
		conditionString + ` AND (LOWER(def."name") LIKE $1)`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(parameter.Search)+"%").Scan(&count)

	return data, count, err
}

func (repository MuUserRepository) Add(c context.Context, model *models.MuUser) (res *string, err error) {

	statement := `INSERT INTO mu_user 
	(	id_branch,id_facebook,id_google,id_apple,name,
		username,password,gender,qr_code,email_user,
		no_telp,address_user,level,birthdate,birthplace,
		referral_code,nik_user,img_ktp,verif_status_ktp,
		created_at_user,created_by_user,active_user,role_group_id
	)
	
	VALUES ( $1, $2, $3, $4, $5,$6, $7, $8, $9 ,$10,
		$11, $12, $13, $14, $15, $16, $17, $18, $19 , $20, $21, $22, $23 ) RETURNING id_user`

	err = repository.DB.QueryRowContext(c, statement,
		str.EmptyString(*model.BranchID), model.FbId, model.GoogleId, model.AppleId, model.Name,
		model.UserName, model.Password, model.Gender, model.QrCode, model.Email,
		model.NoTelp, model.Address, model.Level, model.BirthDate, model.BirthDatePlace,
		model.ReferalCode, model.NIK, model.ImgKTP, model.VerifStatusKtp,
		model.CreatedAt, model.CreatedBy, model.UserActive, model.RoleGroupId).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository MuUserRepository) Edit(c context.Context, model *models.MuUser) (res string, err error) {
	statement := `UPDATE mu_user SET 
	
	id_branch = $1, id_facebook = $2, id_google = $3, id_apple = $4, name = $5, 
	username = $6, gender = $7,  email_user = $8, no_telp = $9, 
	address_user = $10, level = $11, birthdate = $12, birthplace = $13,	nik_user = $14, 
	img_ktp = $15, verif_status_ktp = $16,  updated_at_user = $17, 
	updated_by_user = $18, active_user = $19, role_group_id = $20 , img_profile = $21
	WHERE id_user = $22 RETURNING id_user`
	err = repository.DB.QueryRowContext(c, statement,
		str.EmptyString(*model.BranchID), model.FbId, model.GoogleId, model.AppleId, model.Name,
		model.UserName, model.Gender, model.Email, model.NoTelp,
		model.Address, model.Level, model.BirthDate, model.BirthDatePlace, model.NIK,
		model.ImgKTP, model.VerifStatusKtp, model.UpdatedAt,
		model.UpdatedBy, model.UserActive, model.RoleGroupId, model.ImgProfile, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (repository MuUserRepository) Delete(c context.Context, id string, user_id string, now time.Time) (res string, err error) {
	statement := `UPDATE mu_user SET updated_at_user = $1, deleted_at_user = $2,
	deleted_by_user = $3,updated_by_user = $4
	 WHERE id_user = $5 RETURNING id_user`

	err = repository.DB.QueryRowContext(c, statement, now, now, user_id, user_id, id).Scan(&res)

	if err != nil {
		return res, err
	}
	return res, err
}

func (repository MuUserRepository) SetActiveUser(c context.Context, model *models.MuUser) (res string, err error) {
	statement := `UPDATE mu_user SET active_user = $1, updated_at_user = $2,updated_by_user = $3 WHERE id_user = $4 RETURNING id_user`

	err = repository.DB.QueryRowContext(c, statement, model.UserActive, model.UpdatedAt, model.UpdatedBy, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (repository MuUserRepository) UpdatePassword(c context.Context, model *models.MuUser) (res string, err error) {
	statement := `UPDATE mu_user SET 
	
	password = $1,updated_at_user = $2, 
	updated_by_user = $3
	WHERE id_user = $4 RETURNING id_user`
	err = repository.DB.QueryRowContext(c, statement,
		model.Password, model.UpdatedAt, model.UpdatedBy, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}
	return res, err
}
