import React from 'react';

import { FormField } from '../FormField/FormField';
import { InlineField } from '../Forms/InlineField';
import { SecretFormField } from '../SecretFormField/SecretFormField';

import { HttpSettingsProps } from './types';

export const BasicAuthSettings: React.FC<HttpSettingsProps> = ({ dataSourceConfig, onChange }) => {
  const password = dataSourceConfig.secureJsonData ? dataSourceConfig.secureJsonData.basicAuthPassword : '';

  const onPasswordReset = () => {
    onChange({
      ...dataSourceConfig,
      secureJsonData: {
        ...dataSourceConfig.secureJsonData,
        basicAuthPassword: '',
      },
      secureJsonFields: {
        ...dataSourceConfig.secureJsonFields,
        basicAuthPassword: false,
      },
    });
  };

  const onPasswordChange = (event: React.SyntheticEvent<HTMLInputElement>) => {
    onChange({
      ...dataSourceConfig,
      secureJsonData: {
        ...dataSourceConfig.secureJsonData,
        basicAuthPassword: event.currentTarget.value,
      },
    });
  };

  return (
    <>
      <InlineField>
        <FormField
          label="User"
          labelWidth={10}
          inputWidth={18}
          placeholder="user"
          value={dataSourceConfig.basicAuthUser}
          onChange={(event) => onChange({ ...dataSourceConfig, basicAuthUser: event.currentTarget.value })}
        />
      </InlineField>
      <InlineField>
        <SecretFormField
          isConfigured={!!(dataSourceConfig.secureJsonFields && dataSourceConfig.secureJsonFields.basicAuthPassword)}
          value={password || ''}
          inputWidth={18}
          labelWidth={10}
          onReset={onPasswordReset}
          onChange={onPasswordChange}
        />
      </InlineField>
    </>
  );
};
