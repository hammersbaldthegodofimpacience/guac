package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestSource is the resolver for the ingestSource field.
func (r *mutationResolver) IngestSource(ctx context.Context, source model.SourceInputSpec) (*model.Source, error) {
	return r.Backend.IngestSource(ctx, source)
}

// IngestSources is the resolver for the ingestSources field.
func (r *mutationResolver) IngestSources(ctx context.Context, sources []*model.SourceInputSpec) ([]*model.Source, error) {
	return r.Backend.IngestSources(ctx, sources)
}

// Sources is the resolver for the sources field.
func (r *queryResolver) Sources(ctx context.Context, sourceSpec model.SourceSpec) ([]*model.Source, error) {
	return r.Backend.Sources(ctx, &sourceSpec)
}
