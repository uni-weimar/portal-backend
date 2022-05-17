// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hm-edu/pki-service/ent/certificate"
	"github.com/hm-edu/pki-service/ent/domain"
)

// CertificateCreate is the builder for creating a Certificate entity.
type CertificateCreate struct {
	config
	mutation *CertificateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (cc *CertificateCreate) SetCreateTime(t time.Time) *CertificateCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableCreateTime(t *time.Time) *CertificateCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CertificateCreate) SetUpdateTime(t time.Time) *CertificateCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableUpdateTime(t *time.Time) *CertificateCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetSslId sets the "sslId" field.
func (cc *CertificateCreate) SetSslId(i int) *CertificateCreate {
	cc.mutation.SetSslId(i)
	return cc
}

// SetNillableSslId sets the "sslId" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableSslId(i *int) *CertificateCreate {
	if i != nil {
		cc.SetSslId(*i)
	}
	return cc
}

// SetSerial sets the "serial" field.
func (cc *CertificateCreate) SetSerial(s string) *CertificateCreate {
	cc.mutation.SetSerial(s)
	return cc
}

// SetNillableSerial sets the "serial" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableSerial(s *string) *CertificateCreate {
	if s != nil {
		cc.SetSerial(*s)
	}
	return cc
}

// SetCommonName sets the "commonName" field.
func (cc *CertificateCreate) SetCommonName(s string) *CertificateCreate {
	cc.mutation.SetCommonName(s)
	return cc
}

// SetNotBefore sets the "notBefore" field.
func (cc *CertificateCreate) SetNotBefore(t time.Time) *CertificateCreate {
	cc.mutation.SetNotBefore(t)
	return cc
}

// SetNillableNotBefore sets the "notBefore" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableNotBefore(t *time.Time) *CertificateCreate {
	if t != nil {
		cc.SetNotBefore(*t)
	}
	return cc
}

// SetNotAfter sets the "notAfter" field.
func (cc *CertificateCreate) SetNotAfter(t time.Time) *CertificateCreate {
	cc.mutation.SetNotAfter(t)
	return cc
}

// SetNillableNotAfter sets the "notAfter" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableNotAfter(t *time.Time) *CertificateCreate {
	if t != nil {
		cc.SetNotAfter(*t)
	}
	return cc
}

// SetIssuedBy sets the "issuedBy" field.
func (cc *CertificateCreate) SetIssuedBy(s string) *CertificateCreate {
	cc.mutation.SetIssuedBy(s)
	return cc
}

// SetNillableIssuedBy sets the "issuedBy" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableIssuedBy(s *string) *CertificateCreate {
	if s != nil {
		cc.SetIssuedBy(*s)
	}
	return cc
}

// SetCreated sets the "created" field.
func (cc *CertificateCreate) SetCreated(t time.Time) *CertificateCreate {
	cc.mutation.SetCreated(t)
	return cc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableCreated(t *time.Time) *CertificateCreate {
	if t != nil {
		cc.SetCreated(*t)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CertificateCreate) SetStatus(c certificate.Status) *CertificateCreate {
	cc.mutation.SetStatus(c)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CertificateCreate) SetNillableStatus(c *certificate.Status) *CertificateCreate {
	if c != nil {
		cc.SetStatus(*c)
	}
	return cc
}

// AddDomainIDs adds the "domains" edge to the Domain entity by IDs.
func (cc *CertificateCreate) AddDomainIDs(ids ...int) *CertificateCreate {
	cc.mutation.AddDomainIDs(ids...)
	return cc
}

// AddDomains adds the "domains" edges to the Domain entity.
func (cc *CertificateCreate) AddDomains(d ...*Domain) *CertificateCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cc.AddDomainIDs(ids...)
}

// Mutation returns the CertificateMutation object of the builder.
func (cc *CertificateCreate) Mutation() *CertificateMutation {
	return cc.mutation
}

// Save creates the Certificate in the database.
func (cc *CertificateCreate) Save(ctx context.Context) (*Certificate, error) {
	var (
		err  error
		node *Certificate
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CertificateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CertificateCreate) SaveX(ctx context.Context) *Certificate {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CertificateCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CertificateCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CertificateCreate) defaults() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		if certificate.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized certificate.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := certificate.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		if certificate.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized certificate.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := certificate.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := certificate.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CertificateCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Certificate.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Certificate.update_time"`)}
	}
	if _, ok := cc.mutation.CommonName(); !ok {
		return &ValidationError{Name: "commonName", err: errors.New(`ent: missing required field "Certificate.commonName"`)}
	}
	if v, ok := cc.mutation.CommonName(); ok {
		if err := certificate.CommonNameValidator(v); err != nil {
			return &ValidationError{Name: "commonName", err: fmt.Errorf(`ent: validator failed for field "Certificate.commonName": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Certificate.status"`)}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := certificate.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Certificate.status": %w`, err)}
		}
	}
	return nil
}

func (cc *CertificateCreate) sqlSave(ctx context.Context) (*Certificate, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CertificateCreate) createSpec() (*Certificate, *sqlgraph.CreateSpec) {
	var (
		_node = &Certificate{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: certificate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: certificate.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: certificate.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: certificate.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.SslId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: certificate.FieldSslId,
		})
		_node.SslId = value
	}
	if value, ok := cc.mutation.Serial(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certificate.FieldSerial,
		})
		_node.Serial = value
	}
	if value, ok := cc.mutation.CommonName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certificate.FieldCommonName,
		})
		_node.CommonName = value
	}
	if value, ok := cc.mutation.NotBefore(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: certificate.FieldNotBefore,
		})
		_node.NotBefore = &value
	}
	if value, ok := cc.mutation.NotAfter(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: certificate.FieldNotAfter,
		})
		_node.NotAfter = value
	}
	if value, ok := cc.mutation.IssuedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certificate.FieldIssuedBy,
		})
		_node.IssuedBy = &value
	}
	if value, ok := cc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: certificate.FieldCreated,
		})
		_node.Created = &value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: certificate.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := cc.mutation.DomainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   certificate.DomainsTable,
			Columns: certificate.DomainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: domain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Certificate.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CertificateUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CertificateCreate) OnConflict(opts ...sql.ConflictOption) *CertificateUpsertOne {
	cc.conflict = opts
	return &CertificateUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Certificate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CertificateCreate) OnConflictColumns(columns ...string) *CertificateUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CertificateUpsertOne{
		create: cc,
	}
}

type (
	// CertificateUpsertOne is the builder for "upsert"-ing
	//  one Certificate node.
	CertificateUpsertOne struct {
		create *CertificateCreate
	}

	// CertificateUpsert is the "OnConflict" setter.
	CertificateUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *CertificateUpsert) SetCreateTime(v time.Time) *CertificateUpsert {
	u.Set(certificate.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateCreateTime() *CertificateUpsert {
	u.SetExcluded(certificate.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CertificateUpsert) SetUpdateTime(v time.Time) *CertificateUpsert {
	u.Set(certificate.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateUpdateTime() *CertificateUpsert {
	u.SetExcluded(certificate.FieldUpdateTime)
	return u
}

// SetSslId sets the "sslId" field.
func (u *CertificateUpsert) SetSslId(v int) *CertificateUpsert {
	u.Set(certificate.FieldSslId, v)
	return u
}

// UpdateSslId sets the "sslId" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateSslId() *CertificateUpsert {
	u.SetExcluded(certificate.FieldSslId)
	return u
}

// AddSslId adds v to the "sslId" field.
func (u *CertificateUpsert) AddSslId(v int) *CertificateUpsert {
	u.Add(certificate.FieldSslId, v)
	return u
}

// ClearSslId clears the value of the "sslId" field.
func (u *CertificateUpsert) ClearSslId() *CertificateUpsert {
	u.SetNull(certificate.FieldSslId)
	return u
}

// SetSerial sets the "serial" field.
func (u *CertificateUpsert) SetSerial(v string) *CertificateUpsert {
	u.Set(certificate.FieldSerial, v)
	return u
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateSerial() *CertificateUpsert {
	u.SetExcluded(certificate.FieldSerial)
	return u
}

// ClearSerial clears the value of the "serial" field.
func (u *CertificateUpsert) ClearSerial() *CertificateUpsert {
	u.SetNull(certificate.FieldSerial)
	return u
}

// SetCommonName sets the "commonName" field.
func (u *CertificateUpsert) SetCommonName(v string) *CertificateUpsert {
	u.Set(certificate.FieldCommonName, v)
	return u
}

// UpdateCommonName sets the "commonName" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateCommonName() *CertificateUpsert {
	u.SetExcluded(certificate.FieldCommonName)
	return u
}

// SetNotBefore sets the "notBefore" field.
func (u *CertificateUpsert) SetNotBefore(v time.Time) *CertificateUpsert {
	u.Set(certificate.FieldNotBefore, v)
	return u
}

// UpdateNotBefore sets the "notBefore" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateNotBefore() *CertificateUpsert {
	u.SetExcluded(certificate.FieldNotBefore)
	return u
}

// ClearNotBefore clears the value of the "notBefore" field.
func (u *CertificateUpsert) ClearNotBefore() *CertificateUpsert {
	u.SetNull(certificate.FieldNotBefore)
	return u
}

// SetNotAfter sets the "notAfter" field.
func (u *CertificateUpsert) SetNotAfter(v time.Time) *CertificateUpsert {
	u.Set(certificate.FieldNotAfter, v)
	return u
}

// UpdateNotAfter sets the "notAfter" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateNotAfter() *CertificateUpsert {
	u.SetExcluded(certificate.FieldNotAfter)
	return u
}

// ClearNotAfter clears the value of the "notAfter" field.
func (u *CertificateUpsert) ClearNotAfter() *CertificateUpsert {
	u.SetNull(certificate.FieldNotAfter)
	return u
}

// SetIssuedBy sets the "issuedBy" field.
func (u *CertificateUpsert) SetIssuedBy(v string) *CertificateUpsert {
	u.Set(certificate.FieldIssuedBy, v)
	return u
}

// UpdateIssuedBy sets the "issuedBy" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateIssuedBy() *CertificateUpsert {
	u.SetExcluded(certificate.FieldIssuedBy)
	return u
}

// ClearIssuedBy clears the value of the "issuedBy" field.
func (u *CertificateUpsert) ClearIssuedBy() *CertificateUpsert {
	u.SetNull(certificate.FieldIssuedBy)
	return u
}

// SetCreated sets the "created" field.
func (u *CertificateUpsert) SetCreated(v time.Time) *CertificateUpsert {
	u.Set(certificate.FieldCreated, v)
	return u
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateCreated() *CertificateUpsert {
	u.SetExcluded(certificate.FieldCreated)
	return u
}

// ClearCreated clears the value of the "created" field.
func (u *CertificateUpsert) ClearCreated() *CertificateUpsert {
	u.SetNull(certificate.FieldCreated)
	return u
}

// SetStatus sets the "status" field.
func (u *CertificateUpsert) SetStatus(v certificate.Status) *CertificateUpsert {
	u.Set(certificate.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CertificateUpsert) UpdateStatus() *CertificateUpsert {
	u.SetExcluded(certificate.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Certificate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *CertificateUpsertOne) UpdateNewValues() *CertificateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(certificate.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Certificate.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CertificateUpsertOne) Ignore() *CertificateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CertificateUpsertOne) DoNothing() *CertificateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CertificateCreate.OnConflict
// documentation for more info.
func (u *CertificateUpsertOne) Update(set func(*CertificateUpsert)) *CertificateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CertificateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CertificateUpsertOne) SetCreateTime(v time.Time) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateCreateTime() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CertificateUpsertOne) SetUpdateTime(v time.Time) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateUpdateTime() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetSslId sets the "sslId" field.
func (u *CertificateUpsertOne) SetSslId(v int) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetSslId(v)
	})
}

// AddSslId adds v to the "sslId" field.
func (u *CertificateUpsertOne) AddSslId(v int) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.AddSslId(v)
	})
}

// UpdateSslId sets the "sslId" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateSslId() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateSslId()
	})
}

// ClearSslId clears the value of the "sslId" field.
func (u *CertificateUpsertOne) ClearSslId() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearSslId()
	})
}

// SetSerial sets the "serial" field.
func (u *CertificateUpsertOne) SetSerial(v string) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetSerial(v)
	})
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateSerial() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateSerial()
	})
}

// ClearSerial clears the value of the "serial" field.
func (u *CertificateUpsertOne) ClearSerial() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearSerial()
	})
}

// SetCommonName sets the "commonName" field.
func (u *CertificateUpsertOne) SetCommonName(v string) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetCommonName(v)
	})
}

// UpdateCommonName sets the "commonName" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateCommonName() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateCommonName()
	})
}

// SetNotBefore sets the "notBefore" field.
func (u *CertificateUpsertOne) SetNotBefore(v time.Time) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetNotBefore(v)
	})
}

// UpdateNotBefore sets the "notBefore" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateNotBefore() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateNotBefore()
	})
}

// ClearNotBefore clears the value of the "notBefore" field.
func (u *CertificateUpsertOne) ClearNotBefore() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearNotBefore()
	})
}

// SetNotAfter sets the "notAfter" field.
func (u *CertificateUpsertOne) SetNotAfter(v time.Time) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetNotAfter(v)
	})
}

// UpdateNotAfter sets the "notAfter" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateNotAfter() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateNotAfter()
	})
}

// ClearNotAfter clears the value of the "notAfter" field.
func (u *CertificateUpsertOne) ClearNotAfter() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearNotAfter()
	})
}

// SetIssuedBy sets the "issuedBy" field.
func (u *CertificateUpsertOne) SetIssuedBy(v string) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetIssuedBy(v)
	})
}

// UpdateIssuedBy sets the "issuedBy" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateIssuedBy() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateIssuedBy()
	})
}

// ClearIssuedBy clears the value of the "issuedBy" field.
func (u *CertificateUpsertOne) ClearIssuedBy() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearIssuedBy()
	})
}

// SetCreated sets the "created" field.
func (u *CertificateUpsertOne) SetCreated(v time.Time) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetCreated(v)
	})
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateCreated() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateCreated()
	})
}

// ClearCreated clears the value of the "created" field.
func (u *CertificateUpsertOne) ClearCreated() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearCreated()
	})
}

// SetStatus sets the "status" field.
func (u *CertificateUpsertOne) SetStatus(v certificate.Status) *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CertificateUpsertOne) UpdateStatus() *CertificateUpsertOne {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *CertificateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CertificateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CertificateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CertificateUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CertificateUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CertificateCreateBulk is the builder for creating many Certificate entities in bulk.
type CertificateCreateBulk struct {
	config
	builders []*CertificateCreate
	conflict []sql.ConflictOption
}

// Save creates the Certificate entities in the database.
func (ccb *CertificateCreateBulk) Save(ctx context.Context) ([]*Certificate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Certificate, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CertificateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CertificateCreateBulk) SaveX(ctx context.Context) []*Certificate {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CertificateCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CertificateCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Certificate.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CertificateUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CertificateCreateBulk) OnConflict(opts ...sql.ConflictOption) *CertificateUpsertBulk {
	ccb.conflict = opts
	return &CertificateUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Certificate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CertificateCreateBulk) OnConflictColumns(columns ...string) *CertificateUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CertificateUpsertBulk{
		create: ccb,
	}
}

// CertificateUpsertBulk is the builder for "upsert"-ing
// a bulk of Certificate nodes.
type CertificateUpsertBulk struct {
	create *CertificateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Certificate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *CertificateUpsertBulk) UpdateNewValues() *CertificateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(certificate.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Certificate.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CertificateUpsertBulk) Ignore() *CertificateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CertificateUpsertBulk) DoNothing() *CertificateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CertificateCreateBulk.OnConflict
// documentation for more info.
func (u *CertificateUpsertBulk) Update(set func(*CertificateUpsert)) *CertificateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CertificateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CertificateUpsertBulk) SetCreateTime(v time.Time) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateCreateTime() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CertificateUpsertBulk) SetUpdateTime(v time.Time) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateUpdateTime() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetSslId sets the "sslId" field.
func (u *CertificateUpsertBulk) SetSslId(v int) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetSslId(v)
	})
}

// AddSslId adds v to the "sslId" field.
func (u *CertificateUpsertBulk) AddSslId(v int) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.AddSslId(v)
	})
}

// UpdateSslId sets the "sslId" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateSslId() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateSslId()
	})
}

// ClearSslId clears the value of the "sslId" field.
func (u *CertificateUpsertBulk) ClearSslId() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearSslId()
	})
}

// SetSerial sets the "serial" field.
func (u *CertificateUpsertBulk) SetSerial(v string) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetSerial(v)
	})
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateSerial() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateSerial()
	})
}

// ClearSerial clears the value of the "serial" field.
func (u *CertificateUpsertBulk) ClearSerial() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearSerial()
	})
}

// SetCommonName sets the "commonName" field.
func (u *CertificateUpsertBulk) SetCommonName(v string) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetCommonName(v)
	})
}

// UpdateCommonName sets the "commonName" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateCommonName() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateCommonName()
	})
}

// SetNotBefore sets the "notBefore" field.
func (u *CertificateUpsertBulk) SetNotBefore(v time.Time) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetNotBefore(v)
	})
}

// UpdateNotBefore sets the "notBefore" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateNotBefore() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateNotBefore()
	})
}

// ClearNotBefore clears the value of the "notBefore" field.
func (u *CertificateUpsertBulk) ClearNotBefore() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearNotBefore()
	})
}

// SetNotAfter sets the "notAfter" field.
func (u *CertificateUpsertBulk) SetNotAfter(v time.Time) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetNotAfter(v)
	})
}

// UpdateNotAfter sets the "notAfter" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateNotAfter() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateNotAfter()
	})
}

// ClearNotAfter clears the value of the "notAfter" field.
func (u *CertificateUpsertBulk) ClearNotAfter() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearNotAfter()
	})
}

// SetIssuedBy sets the "issuedBy" field.
func (u *CertificateUpsertBulk) SetIssuedBy(v string) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetIssuedBy(v)
	})
}

// UpdateIssuedBy sets the "issuedBy" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateIssuedBy() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateIssuedBy()
	})
}

// ClearIssuedBy clears the value of the "issuedBy" field.
func (u *CertificateUpsertBulk) ClearIssuedBy() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearIssuedBy()
	})
}

// SetCreated sets the "created" field.
func (u *CertificateUpsertBulk) SetCreated(v time.Time) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetCreated(v)
	})
}

// UpdateCreated sets the "created" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateCreated() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateCreated()
	})
}

// ClearCreated clears the value of the "created" field.
func (u *CertificateUpsertBulk) ClearCreated() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.ClearCreated()
	})
}

// SetStatus sets the "status" field.
func (u *CertificateUpsertBulk) SetStatus(v certificate.Status) *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CertificateUpsertBulk) UpdateStatus() *CertificateUpsertBulk {
	return u.Update(func(s *CertificateUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *CertificateUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CertificateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CertificateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CertificateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
