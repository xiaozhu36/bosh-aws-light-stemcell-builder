#!/bin/bash

set -eu

: ${AWS_ACCESS_KEY_ID:?}
: ${AWS_SECRET_ACCESS_KEY:?}
: ${AWS_DEFAULT_REGION:?}
: ${OUTPUT_BUCKET:?}
: ${BOSHIO_TOKEN:=""}

# inputs
light_stemcell_dir="$PWD/light-stemcell"

echo "************ 14: $light_stemcell_dir"

light_stemcell_path="$(echo ${light_stemcell_dir}/*.tgz)"
echo "************ 17: $light_stemcell_path"
light_stemcell_name="$(basename "${light_stemcell_path}")"
echo "************ 19: $light_stemcell_name"

tar -Oxf ${light_stemcell_path} stemcell.MF > /tmp/stemcell.MF

OS_NAME="$(bosh int /tmp/stemcell.MF --path /operating_system)"
STEMCELL_VERSION="$(bosh int /tmp/stemcell.MF --path /version)"

git clone stemcells-index stemcells-index-output

meta4_folder=$PWD/stemcells-index-output/published/$OS_NAME/$STEMCELL_VERSION
meta4_path=${meta4_folder}/stemcells.alicloud.meta4

mkdir -p "$(dirname "${meta4_path}")"
meta4 create --metalink="$meta4_path"

meta4 import-file --metalink="$meta4_path" --version="$STEMCELL_VERSION" "light-stemcell/${light_stemcell_name}"
meta4 file-set-url --metalink="$meta4_path" --file="${light_stemcell_name}" "https://s3.amazonaws.com/${OUTPUT_BUCKET}/${light_stemcell_name}"

pushd stemcells-index-output > /dev/null
  git add -A
  git -c user.email="ci@localhost" -c user.name="CI Bot" \
    commit -m "publish: $OS_NAME/$STEMCELL_VERSION"
popd > /dev/null

echo "Uploading light stemcell ${light_stemcell_name} to ${OUTPUT_BUCKET}..."
#aws s3 cp "${light_stemcell_path}" "s3://${OUTPUT_BUCKET}"
wget -q http://aliyun-cli.oss-cn-hangzhou.aliyuncs.com/aliyun-cli-linux-3.0.4-amd64.tgz
tar -xzf aliyun-cli-linux-3.0.4-amd64.tgz -C /usr/bin

aliyun oss cp "${light_stemcell_path}" "oss://${OUTPUT_BUCKET}/${light_stemcell_name}" --access-key-id ${AWS_ACCESS_KEY_ID} --access-key-secret ${AWS_SECRET_ACCESS_KEY} --region ${AWS_DEFAULT_REGION}

echo "Stemcell metalink"
cat "$meta4_path"
