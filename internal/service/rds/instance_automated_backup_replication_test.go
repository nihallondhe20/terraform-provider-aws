package rds_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func TestAccRDSInstanceAutomatedBackupReplication_basic(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_db_instance_automated_backup_replication.test"

	var providers []*schema.Provider

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, rds.EndpointsID),
		ProviderFactories: acctest.FactoriesAlternate(&providers),
		CheckDestroy:      testAccCheckInstanceRoleAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInstanceAutomatedBackupReplicationConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "retention_period", "7"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRDSInstanceAutomatedBackupReplication_retentionPeriod(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_db_instance_automated_backup_replication.test"

	var providers []*schema.Provider

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, rds.EndpointsID),
		ProviderFactories: acctest.FactoriesAlternate(&providers),
		CheckDestroy:      testAccCheckInstanceRoleAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInstanceAutomatedBackupReplicationConfig_retentionPeriod(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "retention_period", "14"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRDSInstanceAutomatedBackupReplication_kmsEncrypted(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_db_instance_automated_backup_replication.test"

	var providers []*schema.Provider

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, rds.EndpointsID),
		ProviderFactories: acctest.FactoriesAlternate(&providers),
		CheckDestroy:      testAccCheckInstanceRoleAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInstanceAutomatedBackupReplicationConfig_kmsEncrypted(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "retention_period", "7"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccInstanceAutomatedBackupReplicationConfig(rName string) string {
	return acctest.ConfigCompose(
		acctest.ConfigMultipleRegionProvider(2),
		fmt.Sprintf(`
resource "aws_db_instance" "test" {
  allocated_storage       = 10
  identifier              = %[1]q
  engine                  = "postgres"
  engine_version          = "13.4"
  instance_class          = "db.t3.micro"
  name                    = "mydb"
  username                = "masterusername"
  password                = "mustbeeightcharacters"
  backup_retention_period = 7
  skip_final_snapshot     = true
}

resource "aws_db_instance_automated_backup_replication" "test" {
  source_db_instance_arn = aws_db_instance.test.arn

  provider = "awsalternate"
}
`, rName))
}

func testAccInstanceAutomatedBackupReplicationConfig_retentionPeriod(rName string) string {
	return acctest.ConfigCompose(
		acctest.ConfigMultipleRegionProvider(2),
		fmt.Sprintf(`
resource "aws_db_instance" "test" {
  allocated_storage       = 10
  identifier              = %[1]q
  engine                  = "postgres"
  engine_version          = "13.4"
  instance_class          = "db.t3.micro"
  name                    = "mydb"
  username                = "masterusername"
  password                = "mustbeeightcharacters"
  backup_retention_period = 7
  skip_final_snapshot     = true
}

resource "aws_db_instance_automated_backup_replication" "test" {
  source_db_instance_arn = aws_db_instance.test.arn
  retention_period       = 14

  provider = "awsalternate"
}
`, rName))
}

func testAccInstanceAutomatedBackupReplicationConfig_kmsEncrypted(rName string) string {
	return acctest.ConfigCompose(
		acctest.ConfigMultipleRegionProvider(2),
		fmt.Sprintf(`
resource "aws_kms_key" "test" {
  description = %[1]q
  provider    = "awsalternate"
}
		  
resource "aws_db_instance" "test" {
  allocated_storage       = 10
  identifier              = %[1]q
  engine                  = "postgres"
  engine_version          = "13.4"
  instance_class          = "db.t3.micro"
  name                    = "mydb"
  username                = "masterusername"
  password                = "mustbeeightcharacters"
  backup_retention_period = 7
  storage_encrypted       = true
  skip_final_snapshot     = true
}

resource "aws_db_instance_automated_backup_replication" "test" {
  source_db_instance_arn = aws_db_instance.test.arn
  kms_key_id             = aws_kms_key.test.arn

  provider = "awsalternate"
}
`, rName))
}

func testAccCheckInstanceAutomatedBackupReplicationDestroy(s *terraform.State) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).RDSConn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_db_instance_automated_backup_replication" {
			continue
		}

		input := &rds.DescribeDBInstanceAutomatedBackupsInput{
			DBInstanceAutomatedBackupsArn: &rs.Primary.ID,
		}

		output, err := conn.DescribeDBInstanceAutomatedBackups(input)

		if tfawserr.ErrCodeEquals(err, rds.ErrCodeDBInstanceAutomatedBackupNotFoundFault) {
			continue
		}

		if err != nil {
			return err
		}

		if output == nil {
			continue
		}

		return fmt.Errorf("RDS instance backup replication %q still exists", rs.Primary.ID)
	}

	return nil
}
