package server

import (
	"github.com/synchthia/systera-api/database"
	"github.com/synchthia/systera-api/stream"
	pb "github.com/synchthia/systera-api/systerapb"
	"golang.org/x/net/context"
)

func (s *grpcServer) FetchGroups(ctx context.Context, e *pb.FetchGroupsRequest) (*pb.FetchGroupsResponse, error) {
	groups, err := s.mysql.GetAllGroup()

	var allGroups []*pb.GroupEntry
	for _, group := range groups {
		allGroups = append(allGroups, group.ToProtobuf())
	}

	return &pb.FetchGroupsResponse{Groups: allGroups}, err
}

func (s *grpcServer) CreateGroup(ctx context.Context, e *pb.CreateGroupRequest) (*pb.Empty, error) {
	d := database.Groups{}
	d.Name = e.GroupEntry.GroupName
	d.Prefix = e.GroupEntry.GroupPrefix

	var dbPerms []database.Permissions
	for _, v := range e.GroupEntry.Permissions {
		for _, p := range v.Permissions {
			perm := database.Permissions{
				ServerName: v.ServerName,
				Permission: p,
			}
			dbPerms = append(dbPerms, perm)
		}
	}
	d.Permissions = dbPerms

	err := s.mysql.CreateGroup(d)
	if err != nil {
		return &pb.Empty{}, err
	}

	stream.PublishGroup(d.ToProtobuf())

	return &pb.Empty{}, err
}

func (s *grpcServer) RemoveGroup(ctx context.Context, e *pb.RemoveGroupRequest) (*pb.Empty, error) {
	err := s.mysql.RemoveGroup(e.GroupName)
	return &pb.Empty{}, err
}

func (s *grpcServer) UpdateGroup(ctx context.Context, e *pb.UpdateGroupRequest) (*pb.Empty, error) {
	d := database.Groups{}
	d.Name = e.GroupEntry.GroupName
	d.Prefix = e.GroupEntry.GroupPrefix

	var dbPerms []database.Permissions
	for _, v := range e.GroupEntry.Permissions {
		for _, p := range v.Permissions {
			perm := database.Permissions{
				ServerName: v.ServerName,
				Permission: p,
			}
			dbPerms = append(dbPerms, perm)
		}
	}
	d.Permissions = dbPerms

	err := s.mysql.UpdateGroup(d)
	if err != nil {
		return &pb.Empty{}, err
	}

	stream.PublishGroup(d.ToProtobuf())

	return &pb.Empty{}, err
}

func (s *grpcServer) AddPermission(ctx context.Context, e *pb.AddPermissionRequest) (*pb.Empty, error) {
	if err := s.mysql.AddPermission(e.GroupName, e.Target, e.Permissions); err != nil {
		return &pb.Empty{}, err
	}
	data, err := s.mysql.GetGroupData(e.GroupName)
	stream.PublishPerms(e.Target, data.ToProtobuf())

	return &pb.Empty{}, err
}

func (s *grpcServer) RemovePermission(ctx context.Context, e *pb.RemovePermissionRequest) (*pb.Empty, error) {
	if err := s.mysql.RemovePermission(e.GroupName, e.Target, e.Permissions); err != nil {
		return &pb.Empty{}, err
	}
	data, err := s.mysql.GetGroupData(e.GroupName)
	stream.PublishPerms(e.Target, data.ToProtobuf())

	return &pb.Empty{}, err
}
