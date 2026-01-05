import os
import argparse
from infra_terraform import TerraformManager

def parse_arguments():
    parser = argparse.ArgumentParser(description='Infra Terraform CLI Tool')
    parser.add_argument('--action', type=str, choices=['init', 'plan', 'apply', 'destroy'], required=True)
    parser.add_argument('--env', type=str, choices=['dev', 'stg', 'prod'], default='dev')
    return parser.parse_args()

def main():
    args = parse_arguments()
    terraform_manager = TerraformManager(args.env)
    
    if args.action == 'init':
        terraform_manager.init()
    elif args.action == 'plan':
        terraform_manager.plan()
    elif args.action == 'apply':
        terraform_manager.apply()
    elif args.action == 'destroy':
        terraform_manager.destroy()

if __name__ == '__main__':
    main()